package svc

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

// companies 端點的 orderBy 是 enum，只接受這兩個值。
const (
	COMPANY_ORDER_BY_UEN  = "uen"
	COMPANY_ORDER_BY_NAME = "name"
)

// CompanyQuery 是 GET /v2/companies 的查詢條件。Name 為模糊比對。
type CompanyQuery struct {
	Name               string
	OrderBy            string
	OrderDirection     string
	ResponsiveEmployer bool
	Page               Page
}

// CompanyList 是 /v2/companies 的回應，Total 為全站雇主數。
type CompanyList struct {
	Total   int       `json:"total"`
	Results []Company `json:"results"`
	Links   Links     `json:"_links"`
}

// Companies 依名稱查詢雇主，用於把 JD 上的公司名稱解析成 UEN。
func (c *Client) Companies(ctx context.Context, q CompanyQuery) (*CompanyList, error) {
	query, err := pageQuery(q.Page)
	if err != nil {
		return nil, err
	}
	if name := strings.TrimSpace(q.Name); name != "" {
		query.Set("name", name)
	}
	if order := strings.TrimSpace(q.OrderBy); order != "" {
		query.Set("orderBy", order)
	}
	if dir := strings.TrimSpace(q.OrderDirection); dir != "" {
		query.Set("orderDirection", dir)
	}
	if q.ResponsiveEmployer {
		query.Set("responsiveEmployer", "true")
	}

	var list CompanyList
	if err := c.get(ctx, "/v2/companies", query, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Company 依 UEN 取回單一雇主。回應中的 _links 另指向 jobs / addresses / schemes
// 三條子資源，但那三條需要授權，未帶憑證會回 401 —— 公開路徑上沒有「列出某雇主的所有職缺」，
// 只能以 Search 的關鍵字比對雇主名稱近似取得。
func (c *Client) Company(ctx context.Context, uen string) (*Company, error) {
	id := strings.TrimSpace(uen)
	if id == "" {
		return nil, fmt.Errorf("company uen must not be empty")
	}

	var company Company
	if err := c.get(ctx, "/v2/companies/"+url.PathEscape(id), nil, &company); err != nil {
		return nil, err
	}
	return &company, nil
}
