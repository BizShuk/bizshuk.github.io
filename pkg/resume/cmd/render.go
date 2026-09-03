package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/bizshuk/gosdk/tui"

	"github.com/BizShuk/bizshuk.github.io/pkg/resume/svc"
)

// jsonFlag 由每個命令各自註冊，指向同一份輸出模式決策。
var jsonFlag bool

// tagPattern 清掉 JD 的 HTML 標記。API 的 description 是 employer 貼上的 rich text，
// 直接印到終端機會被標籤淹沒。
var tagPattern = regexp.MustCompile(`<[^>]*>`)

// writeJSON 以縮排 JSON 輸出，供管線接 jq 或直接落檔。
func writeJSON(w io.Writer, v any) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	return encoder.Encode(v)
}

// writeJobTable 印出職缺摘要表。薪酬缺漏時留白，不以 0 冒充。
func writeJobTable(w io.Writer, jobs []svc.Job) {
	table := &tui.Table{
		Headers: []string{"UUID", "TITLE", "COMPANY", "LEVEL", "SALARY", "POSTED"},
		Align:   []int{0, 0, 0, 0, 1, 0},
	}
	for _, job := range jobs {
		table.Rows = append(table.Rows, []tui.Cell{
			job.UUID,
			truncate(job.Title, 52),
			truncate(companyName(&job), 32),
			positionLevels(&job),
			salaryRange(job.Salary),
			postedDate(&job),
		})
	}
	table.Draw(w, false, false)
}

// writeCompanyTable 印出雇主摘要表。
func writeCompanyTable(w io.Writer, companies []svc.Company) {
	table := &tui.Table{
		Headers: []string{"UEN", "NAME", "INDUSTRY", "URL"},
		Align:   []int{0, 0, 0, 0},
	}
	for _, company := range companies {
		table.Rows = append(table.Rows, []tui.Cell{
			company.UEN,
			truncate(company.Name, 40),
			truncate(company.SSICDescription2020, 34),
			truncate(company.CompanyURL, 34),
		})
	}
	table.Draw(w, false, false)
}

// writeJobDetail 印出單一職缺的可讀版本，JD 全文去標籤後附在最後。
func writeJobDetail(w io.Writer, job *svc.Job) {
	fmt.Fprintf(w, "%s\n", job.Title)
	fmt.Fprintf(w, "company : %s\n", companyName(job))
	fmt.Fprintf(w, "level   : %s\n", positionLevels(job))
	fmt.Fprintf(w, "salary  : %s\n", salaryRange(job.Salary))
	fmt.Fprintf(w, "exp     : %d years minimum\n", job.MinimumYearsExperience)
	fmt.Fprintf(w, "skills  : %s\n", skillNames(job))
	fmt.Fprintf(w, "posted  : %s\n", postedDate(job))
	if job.Metadata != nil {
		fmt.Fprintf(w, "url     : %s\n", job.Metadata.JobDetailsURL)
	}
	fmt.Fprintf(w, "\n%s\n", plainText(job.Description))
}

// plainText 把 JD 的 HTML 轉成可讀純文字：區塊標籤換行，其餘標籤直接移除。
func plainText(html string) string {
	text := strings.NewReplacer(
		"</p>", "\n", "</li>", "\n", "</div>", "\n",
		"<br>", "\n", "<br/>", "\n", "<br />", "\n",
		"<li>", "- ",
	).Replace(html)
	text = tagPattern.ReplaceAllString(text, "")
	text = strings.NewReplacer("&amp;", "&", "&lt;", "<", "&gt;", ">", "&nbsp;", " ", "&quot;", `"`, "&#39;", "'").Replace(text)

	var lines []string
	for _, line := range strings.Split(text, "\n") {
		if trimmed := strings.TrimSpace(line); trimmed != "" {
			lines = append(lines, trimmed)
		}
	}
	return strings.Join(lines, "\n")
}

// companyName 優先取實際招募的公司；獵頭代發的職缺 postedCompany 才是仲介。
func companyName(job *svc.Job) string {
	if job.HiringCompany != nil && job.HiringCompany.Name != "" {
		return job.HiringCompany.Name
	}
	if job.PostedCompany != nil {
		return job.PostedCompany.Name
	}
	return ""
}

// salaryRange 格式化薪酬區間，缺值時回空字串而非 0。
func salaryRange(salary *svc.Salary) string {
	if salary == nil || (salary.Minimum == 0 && salary.Maximum == 0) {
		return ""
	}
	period := ""
	if salary.Type != nil {
		period = " " + salary.Type.SalaryType
	}
	return fmt.Sprintf("%d-%d%s", salary.Minimum, salary.Maximum, period)
}

// positionLevels 把職級陣列攤平成逗號分隔字串。
func positionLevels(job *svc.Job) string {
	names := make([]string, 0, len(job.PositionLevels))
	for _, level := range job.PositionLevels {
		names = append(names, level.Position)
	}
	return strings.Join(names, ", ")
}

// skillNames 把技能標籤攤平成逗號分隔字串。
func skillNames(job *svc.Job) string {
	names := make([]string, 0, len(job.Skills))
	for _, skill := range job.Skills {
		names = append(names, skill.Skill)
	}
	return strings.Join(names, ", ")
}

// postedDate 取張貼日；缺 metadata 時回空字串。
func postedDate(job *svc.Job) string {
	if job.Metadata == nil {
		return ""
	}
	if job.Metadata.NewPostingDate != "" {
		return job.Metadata.NewPostingDate
	}
	return job.Metadata.OriginalPostingDate
}

// truncate 以 rune 為單位裁切，避免切壞多位元組字元。
func truncate(text string, max int) string {
	runes := []rune(text)
	if len(runes) <= max {
		return text
	}
	if max <= 1 {
		return string(runes[:max])
	}
	return string(runes[:max-1]) + "…"
}
