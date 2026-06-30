package goshopee

import (
	"net/http"
	"net/url"
)

type Option[T any] func(*Client[T])

type DefaultOption = Option[any]

func WithHTTPClient[T any](client *http.Client) Option[T] {
	return func(c *Client[T]) {
		c.Client = client
	}
}

func WithRetry[T any](retries int) Option[T] {
	return func(c *Client[T]) {
		c.retries = retries
	}
}

func WithLogger[T any](logger LeveledLoggerInterface) Option[T] {
	return func(c *Client[T]) {
		c.log = logger
	}
}

func WithProxy[T any](proxyHost string) Option[T] {
	return func(c *Client[T]) {
		proxyURL, err := url.Parse(proxyHost)
		if err != nil {
			return
		}
		c.Client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	}
}

func WithRefreshToken[T any](refreshToken string) Option[T] {
	return func(c *Client[T]) {
		c.RefreshToken = refreshToken
	}
}

func WithOnTokenRefresh[T any](fn func(res *RefreshAccessTokenResponse, meta T)) Option[T] {
	return func(c *Client[T]) {
		c.OnTokenRefresh = fn
	}
}

func WithMeta[T any](meta T) Option[T] {
	return func(c *Client[T]) {
		c.Meta = meta
	}
}

func WithHTTPClientDefault(client *http.Client) DefaultOption {
	return WithHTTPClient[any](client)
}

func WithRetryDefault(retries int) DefaultOption {
	return WithRetry[any](retries)
}

func WithLoggerDefault(logger LeveledLoggerInterface) DefaultOption {
	return WithLogger[any](logger)
}

func WithProxyDefault(proxyHost string) DefaultOption {
	return WithProxy[any](proxyHost)
}

func WithRefreshTokenDefault(refreshToken string) DefaultOption {
	return WithRefreshToken[any](refreshToken)
}

func WithOnTokenRefreshDefault(fn func(res *RefreshAccessTokenResponse, meta any)) DefaultOption {
	return WithOnTokenRefresh[any](fn)
}

func WithMetaDefault(meta any) DefaultOption {
	return WithMeta[any](meta)
}
