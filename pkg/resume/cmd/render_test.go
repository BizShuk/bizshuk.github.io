package cmd

import (
	"strings"
	"testing"

	"github.com/BizShuk/bizshuk.github.io/pkg/resume/svc"
)

// TestPlainTextUnwrapsJobDescription JD 是 employer 貼上的 rich text，
// 區塊標籤要變成換行，清單項要留住項目符號，實體要還原。
func TestPlainTextUnwrapsJobDescription(t *testing.T) {
	html := `<p><strong>Role</strong></p><ul><li>Lead the team</li><li>Own R&amp;D</li></ul><p>Apply now</p>`
	want := "Role\n- Lead the team\n- Own R&D\nApply now"

	if got := plainText(html); got != want {
		t.Errorf("plainText =\n%q\nwant\n%q", got, want)
	}
}

// TestSalaryRangeLeavesMissingValuesBlank 缺薪酬時要留白，不能用 0 冒充成「月薪零元」。
func TestSalaryRangeLeavesMissingValuesBlank(t *testing.T) {
	if got := salaryRange(nil); got != "" {
		t.Errorf("nil salary = %q, want empty", got)
	}
	if got := salaryRange(&svc.Salary{}); got != "" {
		t.Errorf("zero salary = %q, want empty", got)
	}

	salary := &svc.Salary{Minimum: 8000, Maximum: 12000, Type: &svc.SalaryType{SalaryType: "Monthly"}}
	if got := salaryRange(salary); got != "8000-12000 Monthly" {
		t.Errorf("salary = %q, want 8000-12000 Monthly", got)
	}
}

// TestCompanyNamePrefersHiringCompany 獵頭代發的職缺 postedCompany 是仲介，
// 實際雇主在 hiringCompany，履歷比對要看的是後者。
func TestCompanyNamePrefersHiringCompany(t *testing.T) {
	job := svc.Job{
		PostedCompany: &svc.Company{Name: "RECRUITER PTE. LTD."},
		HiringCompany: &svc.Company{Name: "ACTUAL EMPLOYER"},
	}
	if got := companyName(&job); got != "ACTUAL EMPLOYER" {
		t.Errorf("companyName = %q, want ACTUAL EMPLOYER", got)
	}

	job.HiringCompany = nil
	if got := companyName(&job); got != "RECRUITER PTE. LTD." {
		t.Errorf("companyName = %q, want the posting company as fallback", got)
	}

	if got := companyName(&svc.Job{}); got != "" {
		t.Errorf("companyName = %q, want empty when neither is present", got)
	}
}

// TestTruncateCountsRunes 中文職稱以位元組裁切會切壞字元。
func TestTruncateCountsRunes(t *testing.T) {
	if got := truncate("short", 10); got != "short" {
		t.Errorf("truncate = %q, want the input untouched", got)
	}

	got := truncate("軟體工程部門主管職缺", 5)
	if runes := []rune(got); len(runes) != 5 {
		t.Errorf("truncate = %q (%d runes), want 5 runes", got, len(runes))
	}
	if !strings.HasSuffix(got, "…") {
		t.Errorf("truncate = %q, want an ellipsis marking the cut", got)
	}
}
