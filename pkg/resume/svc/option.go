package svc

import (
	"net/http"
	"strings"
	"time"

	gohttp "github.com/bizshuk/gosdk/http"
)

// Option 以函式選項模式調整 Client。
type Option func(*Client)

// WithBaseURL 覆寫 API 根位址，供測試替身或反向代理使用。空字串會被忽略。
func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		if v := strings.TrimSpace(baseURL); v != "" {
			c.baseURL = v
		}
	}
}

// WithUserAgent 覆寫 User-Agent 標頭。空字串會被忽略。
func WithUserAgent(userAgent string) Option {
	return func(c *Client) {
		if v := strings.TrimSpace(userAgent); v != "" {
			c.userAgent = v
		}
	}
}

// WithTimeout 設定單次請求上限。非正值會被忽略。
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		if timeout > 0 {
			c.http.Timeout = timeout
		}
	}
}

// WithRetryPolicy 覆寫重試預算與退避節奏。
func WithRetryPolicy(policy gohttp.RetryPolicy) Option {
	return func(c *Client) { c.retry = policy }
}

// WithHTTPClient 注入自訂的 http.Client，例如需要特定 Transport 或連線池設定時。
func WithHTTPClient(client *http.Client) Option {
	return func(c *Client) {
		if client != nil {
			c.http = client
		}
	}
}
