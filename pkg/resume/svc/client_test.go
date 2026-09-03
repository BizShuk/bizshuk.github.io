package svc

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	gohttp "github.com/bizshuk/gosdk/http"
)

// TestSearchSendsBodyAndQuery 固定住送出的形狀：POST、JSON body、limit/page 走查詢字串。
func TestSearchSendsBodyAndQuery(t *testing.T) {
	var gotMethod, gotPath, gotQuery, gotUserAgent string
	var gotBody SearchRequest

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		gotQuery = r.URL.RawQuery
		gotUserAgent = r.Header.Get("User-Agent")

		raw, _ := io.ReadAll(r.Body)
		_ = json.Unmarshal(raw, &gotBody)

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"total":2,"countWithoutFilters":9,"results":[{"uuid":"a","title":"Engineer"}]}`))
	}))
	defer server.Close()

	client := New(WithBaseURL(server.URL))
	result, err := client.Search(context.Background(), SearchRequest{
		Search:         "engineer",
		Salary:         9000,
		PositionLevels: []string{"Manager"},
	}, Page{Limit: 5, Page: 2})
	if err != nil {
		t.Fatalf("Search: %v", err)
	}

	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want POST", gotMethod)
	}
	if gotPath != "/v2/search" {
		t.Errorf("path = %q, want /v2/search", gotPath)
	}
	if gotQuery != "limit=5&page=2" {
		t.Errorf("query = %q, want limit=5&page=2", gotQuery)
	}
	if gotUserAgent != DEFAULT_USER_AGENT {
		t.Errorf("user agent = %q, want %q", gotUserAgent, DEFAULT_USER_AGENT)
	}
	if gotBody.Search != "engineer" || gotBody.Salary != 9000 || len(gotBody.PositionLevels) != 1 {
		t.Errorf("body = %+v, want the search, salary and level to survive", gotBody)
	}
	if result.Total != 2 || result.CountWithoutFilters != 9 || len(result.Results) != 1 {
		t.Errorf("result = %+v, want 2 total / 9 unfiltered / 1 row", result)
	}
}

// TestJobFetchesDetailPath 確認 detail 走的是帶 uuid 的路徑，且 description 有解出來。
func TestJobFetchesDetailPath(t *testing.T) {
	var gotPath string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		_, _ = w.Write([]byte(`{"uuid":"abc","title":"Lead","description":"<p>Hi</p>","status":{"id":"102","jobStatus":"Open"}}`))
	}))
	defer server.Close()

	job, err := New(WithBaseURL(server.URL)).Job(context.Background(), "abc")
	if err != nil {
		t.Fatalf("Job: %v", err)
	}
	if gotPath != "/v2/jobs/abc" {
		t.Errorf("path = %q, want /v2/jobs/abc", gotPath)
	}
	if job.Description != "<p>Hi</p>" {
		t.Errorf("description = %q, want the raw HTML", job.Description)
	}
	// status.id 在 search 是字串、在 jobs 是數字，FlexInt 要吃得下兩種。
	if job.Status == nil || job.Status.ID != 102 {
		t.Errorf("status = %+v, want id 102", job.Status)
	}
}

// TestJobRejectsEmptyUUID 空 uuid 應在本地擋下，不該浪費一次往返。
func TestJobRejectsEmptyUUID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("client must not call the server for an empty uuid")
	}))
	defer server.Close()

	if _, err := New(WithBaseURL(server.URL)).Job(context.Background(), "  "); err == nil {
		t.Fatal("want an error for an empty uuid")
	}
}

// TestRetryOnServerError 5xx 要重試；重試預算用完才回錯。
func TestRetryOnServerError(t *testing.T) {
	var calls int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		if calls < 3 {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		_, _ = w.Write([]byte(`{"uuid":"abc","title":"Lead"}`))
	}))
	defer server.Close()

	client := New(
		WithBaseURL(server.URL),
		WithRetryPolicy(gohttp.ConstantRetryPolicy(5, time.Millisecond)),
	)
	if _, err := client.Job(context.Background(), "abc"); err != nil {
		t.Fatalf("Job: %v", err)
	}
	if calls != 3 {
		t.Errorf("calls = %d, want 3 (two 502s then a success)", calls)
	}
}

// TestNoRetryOnValidationError 400 是呼叫端自己的參數錯誤，再試也不會變，
// 且錯誤訊息要帶得回逐欄的驗證說明。
func TestNoRetryOnValidationError(t *testing.T) {
	var calls int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"message":"body: bad","status":400,"errors":[{"path":"sortBy.0","message":"must be equal to one of the allowed values","location":"body"}]}`))
	}))
	defer server.Close()

	client := New(
		WithBaseURL(server.URL),
		WithRetryPolicy(gohttp.ConstantRetryPolicy(5, time.Millisecond)),
	)
	_, err := client.Search(context.Background(), SearchRequest{}, Page{Limit: 1})
	if err == nil {
		t.Fatal("want an error for a 400 response")
	}
	if calls != 1 {
		t.Errorf("calls = %d, want 1 (a 4xx must not be retried)", calls)
	}

	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("error type = %T, want *APIError", err)
	}
	if apiErr.Status != http.StatusBadRequest || len(apiErr.Errors) != 1 {
		t.Errorf("error = %+v, want status 400 with one validation item", apiErr)
	}
	if got := apiErr.Error(); !strings.Contains(got, "sortBy.0") {
		t.Errorf("message = %q, want the offending field named", got)
	}
}

// TestPageQueryRejectsOutOfRange 伺服器的 limit 上限在本地就擋掉，錯得比較早也比較清楚。
func TestPageQueryRejectsOutOfRange(t *testing.T) {
	if _, err := pageQuery(Page{Limit: MAX_PAGE_LIMIT + 1}); err == nil {
		t.Error("want an error when limit exceeds the server maximum")
	}
	if _, err := pageQuery(Page{Limit: 10, Page: -1}); err == nil {
		t.Error("want an error for a negative page")
	}

	query, err := pageQuery(Page{})
	if err != nil {
		t.Fatalf("pageQuery: %v", err)
	}
	if query.Get("limit") != "20" || query.Get("page") != "0" {
		t.Errorf("defaults = %v, want limit 20 page 0", query)
	}
}

// TestFlexIntAcceptsBothShapes 同一欄位在不同端點時而字串、時而數字。
func TestFlexIntAcceptsBothShapes(t *testing.T) {
	for _, tc := range []struct {
		raw  string
		want FlexInt
	}{
		{`{"id":102}`, 102},
		{`{"id":"102"}`, 102},
		{`{"id":null}`, 0},
	} {
		var status Status
		if err := json.Unmarshal([]byte(tc.raw), &status); err != nil {
			t.Fatalf("unmarshal %s: %v", tc.raw, err)
		}
		if status.ID != tc.want {
			t.Errorf("%s -> %d, want %d", tc.raw, status.ID, tc.want)
		}
	}
}
