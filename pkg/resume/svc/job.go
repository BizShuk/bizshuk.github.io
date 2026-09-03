package svc

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

// JobList 是 GET /v2/jobs 的回應：不帶關鍵字的全站職缺列表，依張貼時間倒序。
type JobList struct {
	Total   int   `json:"total"`
	Results []Job `json:"results"`
	Links   Links `json:"_links"`
}

// Job 取回單一職缺的完整內容，包含 /v2/search 省略掉的 description。
func (c *Client) Job(ctx context.Context, uuid string) (*Job, error) {
	id := strings.TrimSpace(uuid)
	if id == "" {
		return nil, fmt.Errorf("job uuid must not be empty")
	}

	var job Job
	if err := c.get(ctx, "/v2/jobs/"+id, nil, &job); err != nil {
		return nil, err
	}
	return &job, nil
}

// Jobs 列出全站職缺。minSalary 為月薪下限，0 表示不設限。
// 這條路徑不吃關鍵字，適合做全量抓取；要篩選請改用 Search。
func (c *Client) Jobs(ctx context.Context, page Page, minSalary int) (*JobList, error) {
	query, err := pageQuery(page)
	if err != nil {
		return nil, err
	}
	if minSalary > 0 {
		query.Set("salary", strconv.Itoa(minSalary))
	}

	var list JobList
	if err := c.get(ctx, "/v2/jobs", query, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// ScreeningQuestions 取回職缺的申請篩選題。多數職缺回空陣列。
func (c *Client) ScreeningQuestions(ctx context.Context, uuid string) ([]ScreeningQuestion, error) {
	id := strings.TrimSpace(uuid)
	if id == "" {
		return nil, fmt.Errorf("job uuid must not be empty")
	}

	var questions []ScreeningQuestion
	if err := c.get(ctx, "/v2/jobs/"+id+"/screening-questions", nil, &questions); err != nil {
		return nil, err
	}
	return questions, nil
}
