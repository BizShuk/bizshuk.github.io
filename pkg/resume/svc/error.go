package svc

import (
	"encoding/json"
	"fmt"
	"strings"
)

// APIError 是 MyCareersFuture 回傳的錯誤。它前面掛了一層 OpenAPI request validator，
// 因此參數錯誤會逐欄列在 Errors 裡，錯誤訊息保留這些欄位比只回 status code 有用得多。
type APIError struct {
	Method   string           `json:"-"`
	Endpoint string           `json:"-"`
	Status   int              `json:"status"`
	Message  string           `json:"message"`
	Errors   []ValidationItem `json:"errors"`
	Raw      string           `json:"-"`
}

// ValidationItem 是單一欄位的驗證失敗說明。
type ValidationItem struct {
	Path      string `json:"path"`
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
	Location  string `json:"location"`
}

// newAPIError 盡量把回應解成結構化錯誤；解不開時保留原始內容以免線索消失。
func newAPIError(method, endpoint string, status int, raw []byte) *APIError {
	e := &APIError{Method: method, Endpoint: endpoint, Status: status, Raw: string(raw)}
	if err := json.Unmarshal(raw, e); err != nil || e.Message == "" {
		e.Status = status
		e.Message = strings.TrimSpace(string(raw))
	}
	e.Status = status
	return e
}

// Error 實作 error 介面。
func (e *APIError) Error() string {
	var b strings.Builder
	fmt.Fprintf(&b, "%s %s: %d", e.Method, e.Endpoint, e.Status)
	if e.Message != "" {
		fmt.Fprintf(&b, ": %s", e.Message)
	}
	for _, item := range e.Errors {
		fmt.Fprintf(&b, " [%s %s %s]", item.Location, item.Path, item.Message)
	}
	return b.String()
}
