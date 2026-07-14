package __PACKAGE_NAME__

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/go-querystring/query"
)

const (
	UserAgent = "__PACKAGE_NAME__/1.0.0"

	defaultHttpTimeout = 10
)

type App struct {
	PartnerID   int    `env:"PARTNER_ID"`
	PartnerKey  string `env:"PARTNER_KEY"`
	RedirectURL string `env:"REDIRECT_URL"`
	APIURL      string `env:"API_URL"`
}

type requestAuth struct {
	shopID     uint64
	merchantID uint64
	token      string
}

type authKeyType struct{}

var authKey = authKeyType{}

func authFromContext(ctx context.Context) requestAuth {
	if v, ok := ctx.Value(authKey).(requestAuth); ok {
		return v
	}
	return requestAuth{}
}

type RateLimitInfo struct {
	RequestCount      int
	BucketSize        int
	RetryAfterSeconds float64
}

type Client[T any] struct {
	mu     sync.Mutex
	Client *http.Client
	log    LeveledLoggerInterface

	app App

	baseURL *url.URL

	retries  int
	attempts int

	RateLimits RateLimitInfo

	RefreshToken   string
	OnTokenRefresh func(res *RefreshAccessTokenResponse, meta T)
	Meta           T

	// @SERVICES_SECTION
}

type DefaultClient = Client[any]

func NewClient[T any](app App, opts ...Option[T]) *Client[T] {
	baseURL, err := url.Parse(app.APIURL)
	if err != nil {
		panic(err)
	}

	c := &Client[T]{
		Client:  &http.Client{Timeout: time.Duration(defaultHttpTimeout) * time.Second},
		log:     &LeveledLogger{},
		app:     app,
		baseURL: baseURL,
	}

	// @SERVICES_INIT_SECTION

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func NewDefaultClient(app App, opts ...DefaultOption) *DefaultClient {
	return NewClient(app, opts...)
}

type ResponseError struct {
	Status      int
	Message     string
	Errors      []string
	ShopeeError string
	RequestID   string
}

func (e ResponseError) GetStatus() int    { return e.Status }
func (e ResponseError) GetMessage() string { return e.Message }
func (e ResponseError) GetErrors() []string { return e.Errors }
func (e ResponseError) GetShopeeError() string { return e.ShopeeError }
func (e ResponseError) GetRequestID() string { return e.RequestID }

func (e ResponseError) Error() string {
	msg := e.Message
	if msg == "" && len(e.Errors) > 0 {
		sort.Strings(e.Errors)
		msg = strings.Join(e.Errors, ", ")
	}
	if msg == "" {
		msg = "Unknown Error"
	}
	if e.ShopeeError != "" {
		msg = fmt.Sprintf("shopee: %s [%s]", e.ShopeeError, msg)
	}
	if e.RequestID != "" {
		msg = fmt.Sprintf("%s (RequestID: %s)", msg, e.RequestID)
	}
	return msg
}

func IsShopeeError(err error, shopeeErrCode string) bool {
	if re, ok := err.(ResponseError); ok {
		return re.ShopeeError == shopeeErrCode
	}
	if re, ok := err.(*ResponseError); ok {
		return re.ShopeeError == shopeeErrCode
	}
	return false
}

type ResponseDecodingError struct {
	Body    []byte
	Message string
	Status  int
}

func (e ResponseDecodingError) Error() string { return e.Message }

type RateLimitError struct {
	ResponseError
	RetryAfter int
}

func (c *Client[T]) NewRequest(ctx context.Context, method, relPath string, body, options, headers interface{}, sid uint64, mid uint64, tok string) (*http.Request, error) {
	rel, err := url.Parse(relPath)
	if err != nil {
		return nil, err
	}

	u := c.baseURL.ResolveReference(rel)

	if options != nil {
		optionsQuery, err := query.Values(options)
		if err != nil {
			return nil, err
		}

		jsonListParams := []string{
			"category_id_list",
			"main_item_id",
			"direct_item_id",
			"shop_id_list",
			"enabled_channel_id_list",
		}
		for _, param := range jsonListParams {
			if values, ok := optionsQuery[param]; ok && len(values) > 1 {
				optionsQuery.Set(param, "["+strings.Join(values, ",")+"]")
			} else if ok && len(values) == 1 && strings.Contains(values[0], ",") {
				if !strings.HasPrefix(values[0], "[") {
					optionsQuery.Set(param, "["+values[0]+"]")
				}
			} else if ok && len(values) == 1 {
				if !strings.HasPrefix(values[0], "[") {
					optionsQuery.Set(param, "["+values[0]+"]")
				}
			}
		}

		if values, ok := optionsQuery["item_id_list"]; ok && len(values) > 0 {
			optionsQuery.Set("item_id_list", strings.Join(values, ","))
		}

		for k, values := range u.Query() {
			for _, v := range values {
				optionsQuery.Add(k, v)
			}
		}
		u.RawQuery = optionsQuery.Encode()
	}

	var js []byte = nil
	if body != nil {
		js, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	ctx = context.WithValue(ctx, authKey, requestAuth{shopID: sid, merchantID: mid, token: tok})
	req, err := http.NewRequestWithContext(ctx, method, u.String(), bytes.NewBuffer(js))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", UserAgent)

	c.makeSignature(req, sid, mid, tok)

	return req, nil
}

func (c *Client[T]) WithRefreshToken(tok string) *Client[T] {
	c.RefreshToken = tok
	return c
}

func (c *Client[T]) WithOnTokenRefresh(fn func(res *RefreshAccessTokenResponse, meta T)) *Client[T] {
	c.OnTokenRefresh = fn
	return c
}

func (c *Client[T]) WithMeta(meta T) *Client[T] {
	c.Meta = meta
	return c
}

func (c *Client[T]) makeSignature(req *http.Request, sid uint64, mid uint64, tok string) (string, int64) {
	ts := time.Now().Unix()
	path := req.URL.Path

	var baseStr string
	u := req.URL
	query := u.Query()
	query.Set("partner_id", fmt.Sprintf("%v", c.app.PartnerID))

	isPublicApi := false
	if strings.Contains(path, "/auth/token/get") || strings.Contains(path, "/auth/access_token/get") {
		isPublicApi = true
	}

	if sid != 0 && !isPublicApi {
		baseStr = fmt.Sprintf("%d%s%d%s%d", c.app.PartnerID, path, ts, tok, sid)
		query.Set("shop_id", fmt.Sprintf("%v", sid))
		query.Set("access_token", tok)
	} else if mid != 0 && !isPublicApi {
		baseStr = fmt.Sprintf("%d%s%d%s%d", c.app.PartnerID, path, ts, tok, mid)
		query.Set("merchant_id", fmt.Sprintf("%v", mid))
		query.Set("access_token", tok)
	} else {
		baseStr = fmt.Sprintf("%d%s%d", c.app.PartnerID, path, ts)
	}

	h := hmac.New(sha256.New, []byte(c.app.PartnerKey))
	h.Write([]byte(baseStr))
	result := hex.EncodeToString(h.Sum(nil))

	query.Set("timestamp", fmt.Sprintf("%v", ts))
	query.Set("sign", result)

	u.RawQuery = query.Encode()
	req.URL = u

	return result, ts
}

func (c *Client[T]) doGetHeaders(req *http.Request, v interface{}, skipBody bool) (http.Header, error) {
	var resp *http.Response
	var err error

	retries := c.retries
	refreshAttempts := 0
	c.attempts = 0
	c.logRequest(req, skipBody)

	for {
		c.attempts++

		resp, err = c.Client.Do(req)
		c.logResponse(resp)
		if err != nil {
			return nil, err
		}

		respErr := CheckResponseError(resp)
		if respErr == nil {
			break
		}

		if c.RefreshToken != "" && !strings.Contains(req.URL.Path, "/auth/access_token/get") {
			var shopeeErr string
			if re, ok := respErr.(ResponseError); ok {
				shopeeErr = re.ShopeeError
			} else if re, ok := respErr.(*ResponseError); ok {
				shopeeErr = re.ShopeeError
			}

			if refreshAttempts < c.retries && (shopeeErr == "error_invalid_access_token" || shopeeErr == "error_access_token_expired" || shopeeErr == "invalid_access_token" || shopeeErr == "invalid_acceess_token") {
				a := authFromContext(req.Context())
				refreshRes, err := c.Auth.RefreshAccessToken(req.Context(), a.shopID, a.merchantID, c.RefreshToken)
				if err == nil {
					c.mu.Lock()
					c.RefreshToken = refreshRes.RefreshToken
					if c.OnTokenRefresh != nil {
						c.OnTokenRefresh(refreshRes, c.Meta)
					}
					c.mu.Unlock()
					c.makeSignature(req, a.shopID, a.merchantID, refreshRes.AccessToken)
					resp.Body.Close()
					refreshAttempts++
					continue
				}
			}
		}

		resp.Body.Close()

		if retries <= 1 {
			return nil, respErr
		}

		if rateLimitErr, isRetryErr := respErr.(RateLimitError); isRetryErr {
			wait := time.Duration(rateLimitErr.RetryAfter) * time.Second
			c.log.Debugf("rate limited waiting %s", wait.String())
			time.Sleep(wait)
			retries--
			continue
		}

		var doRetry bool
		switch resp.StatusCode {
		case http.StatusServiceUnavailable:
			c.log.Debugf("service unavailable, retrying")
			doRetry = true
			retries--
		}

		if doRetry {
			continue
		}

		return nil, respErr
	}

	c.logResponse(resp)
	defer resp.Body.Close()

	if v != nil {
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(v)
		if err != nil {
			return nil, err
		}
	}

	return resp.Header, nil
}

func (c *Client[T]) logRequest(req *http.Request, skipBody bool) {
	if req == nil { return }
	if req.URL != nil { c.log.Debugf("%s: %s", req.Method, req.URL.String()) }
	if !skipBody { c.logBody(&req.Body, "SENT: %s") }
}

func (c *Client[T]) logResponse(res *http.Response) {
	if res == nil { return }
	c.log.Debugf("RECV %d: %s", res.StatusCode, res.Status)
	c.logBody(&res.Body, "RESP: %s")
}

func (c *Client[T]) logBody(body *io.ReadCloser, format string) {
	if body == nil || *body == nil { return }
	b, _ := io.ReadAll(*body)
	if len(b) > 0 { c.log.Debugf(format, string(b)) }
	*body = io.NopCloser(bytes.NewBuffer(b))
}

func wrapSpecificError(r *http.Response, err ResponseError) error {
	if err.Status == http.StatusTooManyRequests {
		f, _ := strconv.ParseFloat(r.Header.Get("Retry-After"), 64)
		return RateLimitError{
			ResponseError: err,
			RetryAfter:    int(f),
		}
	}
	if err.Status == http.StatusNotAcceptable {
		err.Message = http.StatusText(err.Status)
	}
	return err
}

func CheckResponseError(r *http.Response) error {
	shopeeError := struct {
		Error     string `json:"error"`
		Message   string `json:"message"`
		RequestID string `json:"request_id"`
	}{}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil { return err }

	defer func() {
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}()

	if len(bodyBytes) > 0 {
		err := json.Unmarshal(bodyBytes, &shopeeError)
		if err != nil {
			if r.StatusCode == http.StatusOK { return nil }
			return ResponseDecodingError{Body: bodyBytes, Message: err.Error(), Status: r.StatusCode}
		}
	}

	if shopeeError.Error == "" && http.StatusOK <= r.StatusCode && r.StatusCode < http.StatusMultipleChoices {
		return nil
	}

	responseError := ResponseError{
		Status:      r.StatusCode,
		Message:     shopeeError.Message,
		ShopeeError: shopeeError.Error,
		RequestID:   shopeeError.RequestID,
	}
	return wrapSpecificError(r, responseError)
}

func (c *Client[T]) CreateAndDo(ctx context.Context, method, relPath string, data, options, headers, resource interface{}, sid uint64, mid uint64, tok string) error {
	_, err := c.createAndDoGetHeaders(ctx, method, relPath, data, options, headers, resource, sid, mid, tok)
	return err
}

func (c *Client[T]) createAndDoGetHeaders(ctx context.Context, method, relPath string, data, options, headers, resource interface{}, sid uint64, mid uint64, tok string) (http.Header, error) {
	if strings.HasPrefix(relPath, "/") {
		relPath = strings.TrimLeft(relPath, "/")
	}
	relPath = path.Join("api/v2", relPath)
	req, err := c.NewRequest(ctx, method, relPath, data, options, headers, sid, mid, tok)
	if err != nil { return nil, err }
	return c.doGetHeaders(req, resource, false)
}

func (c *Client[T]) Get(ctx context.Context, path string, resource, options interface{}, sid uint64, tok string) error {
	return c.CreateAndDo(ctx, "GET", path, nil, options, nil, resource, sid, 0, tok)
}

func (c *Client[T]) Post(ctx context.Context, path string, data, resource interface{}, sid uint64, tok string) error {
	return c.CreateAndDo(ctx, "POST", path, data, nil, nil, resource, sid, 0, tok)
}

func (c *Client[T]) Put(ctx context.Context, path string, data, resource interface{}, sid uint64, tok string) error {
	return c.CreateAndDo(ctx, "PUT", path, data, nil, nil, resource, sid, 0, tok)
}

func (c *Client[T]) Delete(ctx context.Context, path string, sid uint64, tok string) error {
	return c.CreateAndDo(ctx, "DELETE", path, nil, nil, nil, nil, sid, 0, tok)
}

func (c *Client[T]) Upload(ctx context.Context, relPath, fieldname, filename string, resource interface{}, sid uint64, mid uint64, tok string) error {
	req, err := c.NewfileUploadRequest(ctx, relPath, fieldname, filename, sid, mid, tok)
	if err != nil { return err }
	_, err = c.doGetHeaders(req, resource, true)
	return err
}

func (c *Client[T]) UploadFromReader(ctx context.Context, relPath, fieldname, filename string, reader io.Reader, resource interface{}, sid uint64, mid uint64, tok string) error {
	req, err := c.NewUploadFromReaderRequest(ctx, relPath, fieldname, filename, reader, sid, mid, tok)
	if err != nil { return err }
	_, err = c.doGetHeaders(req, resource, true)
	return err
}

func (c *Client[T]) NewfileUploadRequest(ctx context.Context, relPath, paramName, filename string, sid uint64, mid uint64, tok string) (*http.Request, error) {
	if strings.HasPrefix(relPath, "/") {
		relPath = strings.TrimLeft(relPath, "/")
	}
	relPath = path.Join("api/v2", relPath)
	rel, err := url.Parse(relPath)
	if err != nil { return nil, err }
	u := c.baseURL.ResolveReference(rel)
	uri := u.String()

	file, err := os.Open(filename)
	if err != nil { return nil, err }
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(filename))
	if err != nil { return nil, err }
	if _, err = io.Copy(part, file); err != nil { return nil, err }
	err = writer.Close()
	if err != nil { return nil, err }

	ctx = context.WithValue(ctx, authKey, requestAuth{shopID: sid, merchantID: mid, token: tok})
	req, err := http.NewRequestWithContext(ctx, "POST", uri, body)
	if err != nil { return nil, err }
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", UserAgent)
	c.makeSignature(req, sid, mid, tok)

	return req, nil
}

func (c *Client[T]) NewUploadFromReaderRequest(ctx context.Context, relPath, paramName, filename string, reader io.Reader, sid uint64, mid uint64, tok string) (*http.Request, error) {
	if strings.HasPrefix(relPath, "/") {
		relPath = strings.TrimLeft(relPath, "/")
	}
	relPath = path.Join("api/v2", relPath)
	rel, err := url.Parse(relPath)
	if err != nil { return nil, err }
	u := c.baseURL.ResolveReference(rel)
	uri := u.String()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(filename))
	if err != nil { return nil, err }
	if _, err = io.Copy(part, reader); err != nil { return nil, err }
	err = writer.Close()
	if err != nil { return nil, err }

	ctx = context.WithValue(ctx, authKey, requestAuth{shopID: sid, merchantID: mid, token: tok})
	req, err := http.NewRequestWithContext(ctx, "POST", uri, body)
	if err != nil { return nil, err }
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", UserAgent)
	c.makeSignature(req, sid, mid, tok)

	return req, nil
}

type BoolString bool

func (bs *BoolString) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), "\"")
	if strings.ToUpper(s) == "TRUE" { *bs = true; return nil }
	if strings.ToUpper(s) == "FALSE" { *bs = false; return nil }
	var b bool
	if err := json.Unmarshal(data, &b); err != nil { return err }
	*bs = BoolString(b)
	return nil
}

func (bs BoolString) String() string {
	if bs { return "TRUE" }
	return "FALSE"
}
