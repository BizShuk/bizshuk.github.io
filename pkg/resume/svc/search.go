package svc

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// sortBy 允許值。伺服器端是 enum，送別的值會被 OpenAPI validator 擋成 400；
// 留白代表用相關性 (relevancy) 排序，那是預設行為且沒有對應的 enum 字串。
const (
	SORT_NEW_POSTING_DATE   = "new_posting_date"
	SORT_MIN_MONTHLY_SALARY = "min_monthly_salary"
)

// postingCompany 允許值：職缺由雇主自行張貼，或由人力仲介代發。
const (
	POSTED_BY_DIRECT      = "Direct"
	POSTED_BY_THIRD_PARTY = "Third Party"
)

// flexibleWorkArrangements 是以 id 而非名稱篩選的少數欄位之一。
const (
	FWA_FLEXI_HOURS                  = 1
	FWA_TELECOMMUTING                = 2
	FWA_EMPLOYEES_CHOICE_OF_DAYS_OFF = 3
	FWA_STAGGERED_TIME               = 4
	FWA_COMPRESSED_WORK_SCHEDULE     = 5
	FWA_CREATIVE_SCHEDULING          = 6
)

// Page 是所有分頁端點共用的查詢參數，Page 由 0 起算。
type Page struct {
	Limit int
	Page  int
}

// SearchRequest 是 POST /v2/search 的 body。欄位集合由伺服器端的 OpenAPI validator 定義：
// 送出未知欄位不會報錯但也不會生效，因此這裡只列出實際被驗證且實際生效的欄位。
//
// Categories、PositionLevels、EmploymentTypes 與 PostingCompany 都是 enum，
// 值必須與網站上的顯示名稱完全一致 (例如 "Information Technology"、"Full Time")，
// 錯字會換來 400 而不是空結果。
//
// 另外兩個通過驗證但刻意不放進本結構的欄位：
//   - schemes (boolean)：送 true 或 false 的結果數完全相同，在公開路徑上無作用。
//   - jobStatuses ([]int，102 Open / 103 Re-open)：一旦帶上，伺服器改走另一條路徑，
//     關鍵字與其餘所有篩選全部失效，回應也不再帶 countWithoutFilters。
type SearchRequest struct {
	Search                   string   `json:"search,omitempty"`
	SessionID                string   `json:"sessionId,omitempty"`
	Salary                   int      `json:"salary,omitempty"`
	EmploymentTypes          []string `json:"employmentTypes,omitempty"`
	PositionLevels           []string `json:"positionLevels,omitempty"`
	Categories               []string `json:"categories,omitempty"`
	PostingCompany           []string `json:"postingCompany,omitempty"`
	FlexibleWorkArrangements []int    `json:"flexibleWorkArrangements,omitempty"`
	SortBy                   []string `json:"sortBy,omitempty"`
}

// SearchResult 是 /v2/search 的回應。Total 是套用篩選後的筆數，
// CountWithoutFilters 是同一組關鍵字不套篩選的筆數，兩者的差距即篩選的收斂程度。
type SearchResult struct {
	Total               int    `json:"total"`
	CountWithoutFilters int    `json:"countWithoutFilters"`
	SearchRankingID     string `json:"searchRankingId"`
	Results             []Job  `json:"results"`
	Links               Links  `json:"_links"`
}

// Search 查詢職缺。回傳的 Job 不含 description，需要 JD 全文請再呼叫 Job。
func (c *Client) Search(ctx context.Context, req SearchRequest, page Page) (*SearchResult, error) {
	query, err := pageQuery(page)
	if err != nil {
		return nil, err
	}

	var result SearchResult
	if err := c.post(ctx, "/v2/search", query, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SearchAll 反覆翻頁直到取滿 max 筆或伺服器沒有下一頁。
// max <= 0 表示取回符合條件的全部職缺；呼叫端請自行以 context 設上限。
func (c *Client) SearchAll(ctx context.Context, req SearchRequest, max int) ([]Job, error) {
	limit := MAX_PAGE_LIMIT
	if max > 0 && max < limit {
		limit = max
	}

	var jobs []Job
	for page := 0; ; page++ {
		result, err := c.Search(ctx, req, Page{Limit: limit, Page: page})
		if err != nil {
			return jobs, err
		}
		if len(result.Results) == 0 {
			return jobs, nil
		}
		jobs = append(jobs, result.Results...)

		if max > 0 && len(jobs) >= max {
			return jobs[:max], nil
		}
		if result.Links.Next == nil || len(jobs) >= result.Total {
			return jobs, nil
		}
	}
}

// pageQuery 把分頁參數轉成查詢字串，並在本地擋掉伺服器會拒絕的 limit。
func pageQuery(page Page) (url.Values, error) {
	if page.Limit <= 0 {
		page.Limit = 20
	}
	if page.Limit > MAX_PAGE_LIMIT {
		return nil, fmt.Errorf("limit %d exceeds server maximum %d", page.Limit, MAX_PAGE_LIMIT)
	}
	if page.Page < 0 {
		return nil, fmt.Errorf("page %d must not be negative", page.Page)
	}

	query := url.Values{}
	query.Set("limit", strconv.Itoa(page.Limit))
	query.Set("page", strconv.Itoa(page.Page))
	return query, nil
}
