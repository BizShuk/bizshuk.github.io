// Package svc 封裝 MyCareersFuture 公開 API 的讀取端。
//
// 這是 www.mycareersfuture.gov.sg 前端直接呼叫的 backend，非官方對外契約：
// 不需要 API key、不需要 cookie，但也沒有版本承諾，端點與欄位可能隨網站改版而變。
// 因此本套件只做「HTTP 與 JSON 的翻譯」，不對欄位語意做額外假設。
package svc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	gohttp "github.com/bizshuk/gosdk/http"
)

const (
	// DEFAULT_BASE_URL 是 MyCareersFuture 的 API 根位址。
	DEFAULT_BASE_URL = "https://api.mycareersfuture.gov.sg"
	// DEFAULT_TIMEOUT 是單次請求的上限，含連線、傳輸與讀取。
	DEFAULT_TIMEOUT = 30 * time.Second
	// DEFAULT_USER_AGENT 讓對方端能辨識來源；預設 Go client UA 會被部分邊緣節點擋下。
	DEFAULT_USER_AGENT = "bizshuk-resume-jd/1.0 (+https://github.com/BizShuk/bizshuk.github.io)"
	// MAX_PAGE_LIMIT 是伺服器對 limit 查詢參數的硬上限，超過會回 400。
	MAX_PAGE_LIMIT = 100
)

// Client 是所有 MyCareersFuture 呼叫的進入點，可安全並行使用。
type Client struct {
	baseURL   string
	userAgent string
	retry     gohttp.RetryPolicy
	http      *http.Client
}

// New 建立 Client，未提供選項時採用套件預設值。
func New(opts ...Option) *Client {
	c := &Client{
		baseURL:   DEFAULT_BASE_URL,
		userAgent: DEFAULT_USER_AGENT,
		retry:     gohttp.DefaultRetryPolicy(),
		http:      &http.Client{Timeout: DEFAULT_TIMEOUT},
	}
	for _, opt := range opts {
		opt(c)
	}
	c.baseURL = strings.TrimRight(c.baseURL, "/")
	return c
}

// get 發出 GET 並把回應解碼進 out。
func (c *Client) get(ctx context.Context, path string, query url.Values, out any) error {
	return c.do(ctx, http.MethodGet, path, query, nil, out)
}

// post 發出 POST 並把回應解碼進 out。
func (c *Client) post(ctx context.Context, path string, query url.Values, body, out any) error {
	return c.do(ctx, http.MethodPost, path, query, body, out)
}

// do 組出完整 URL 後交給 gosdk 的重試迴圈，成功時解碼 JSON。
// 請求 body 先序列化成位元組再進迴圈，避免每次重試都要重建 io.Reader。
func (c *Client) do(ctx context.Context, method, path string, query url.Values, body, out any) error {
	var payload []byte
	if body != nil {
		encoded, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("marshal request body: %w", err)
		}
		payload = encoded
	}

	endpoint := c.baseURL + path
	if len(query) > 0 {
		endpoint += "?" + query.Encode()
	}

	raw, err := gohttp.Retry(ctx, c.retry, func(ctx context.Context) ([]byte, error) {
		return c.attempt(ctx, method, endpoint, payload)
	})
	if err != nil {
		return err
	}
	if out == nil {
		return nil
	}
	if err := json.Unmarshal(raw, out); err != nil {
		return fmt.Errorf("decode %s %s: %w", method, endpoint, err)
	}
	return nil
}

// attempt 是單次嘗試。連線層錯誤與 429/5xx 標記為可重試，其餘 4xx 立即回傳，
// 因為那是呼叫端自己的參數問題，再試幾次也不會變。
func (c *Client) attempt(ctx context.Context, method, endpoint string, payload []byte) ([]byte, error) {
	var reader io.Reader
	if payload != nil {
		reader = bytes.NewReader(payload)
	}

	req, err := http.NewRequestWithContext(ctx, method, endpoint, reader)
	if err != nil {
		return nil, fmt.Errorf("build request %s %s: %w", method, endpoint, err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.userAgent)
	if payload != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, gohttp.Retryable(fmt.Errorf("call %s %s: %w", method, endpoint, err))
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, gohttp.Retryable(fmt.Errorf("read body of %s %s: %w", method, endpoint, err))
	}

	if resp.StatusCode >= http.StatusBadRequest {
		apiErr := newAPIError(method, endpoint, resp.StatusCode, raw)
		if gohttp.IsRetryableStatus(resp.StatusCode) {
			return nil, gohttp.Retryable(apiErr)
		}
		return nil, apiErr
	}
	return raw, nil
}
