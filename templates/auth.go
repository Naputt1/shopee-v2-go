package __PACKAGE_NAME__

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type AuthService interface {
	GetAuthURL() (string, error)
	GetCancelAuthURL() (string, error)
	GetAccessToken(ctx context.Context, sid uint64, aid uint64, code string) (*AccessTokenResponse, error)
	RefreshAccessToken(ctx context.Context, sid uint64, aid uint64, refresh string) (*RefreshAccessTokenResponse, error)
}

type AccessTokenResponse struct {
	BaseResponse

	AccessToken    string   `json:"access_token"`
	RefreshToken   string   `json:"refresh_token"`
	ExpireIn       int      `json:"expire_in"`
	MerchantIDList []uint64 `json:"merchant_id_list,omitempty"`
	ShopIDList     []uint64 `json:"shop_id_list,omitempty"`
}

type RefreshAccessTokenResponse struct {
	BaseResponse

	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpireIn     int    `json:"expire_in"`
	PartnerID    uint64 `json:"partner_id"`
	MerchantID   uint64 `json:"merchant_id"`
	ShopID       uint64 `json:"shop_id"`
}

type AuthServiceOp[T any] struct {
	client *Client[T]
}

func (s *AuthServiceOp[T]) GetAuthURL() (string, error) {
	return s.authURL("/api/v2/shop/auth_partner")
}

func (s *AuthServiceOp[T]) GetCancelAuthURL() (string, error) {
	return s.authURL("/api/v2/shop/cancel_auth_partner")
}

func (s *AuthServiceOp[T]) authURL(path string) (string, error) {
	rurl := s.client.app.RedirectURL
	ts := time.Now().Unix()
	baseStr := fmt.Sprintf("%d%s%d", s.client.app.PartnerID, path, ts)
	h := hmac.New(sha256.New, []byte(s.client.app.PartnerKey))
	h.Write([]byte(baseStr))
	sign := hex.EncodeToString(h.Sum(nil))
	return fmt.Sprintf("%s%s?partner_id=%d&timestamp=%d&sign=%s&redirect=%s", s.client.app.APIURL, path, s.client.app.PartnerID, ts, sign, rurl), nil
}

func (s *AuthServiceOp[T]) GetAccessToken(ctx context.Context, sid uint64, aid uint64, code string) (*AccessTokenResponse, error) {
	path := "/auth/token/get"
	params := map[string]interface{}{
		"code":       code,
		"partner_id": s.client.app.PartnerID,
	}
	if sid != 0 {
		params["shop_id"] = sid
	} else if aid != 0 {
		params["main_account_id"] = aid
	}
	resp := new(AccessTokenResponse)
	err := s.client.Post(ctx, path, params, resp, 0, "")
	return resp, err
}

func (s *AuthServiceOp[T]) RefreshAccessToken(ctx context.Context, sid uint64, aid uint64, refresh string) (*RefreshAccessTokenResponse, error) {
	path := "/auth/access_token/get"
	params := map[string]interface{}{
		"refresh_token": refresh,
		"partner_id":    s.client.app.PartnerID,
	}
	if sid != 0 {
		params["shop_id"] = sid
	} else if aid != 0 {
		params["main_account_id"] = aid
	}
	resp := new(RefreshAccessTokenResponse)
	err := s.client.Post(ctx, path, params, resp, 0, "")
	return resp, err
}
