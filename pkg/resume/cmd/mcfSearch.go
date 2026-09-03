package cmd

import (
	"context"
	"os"

	"github.com/spf13/cobra"

	"github.com/BizShuk/bizshuk.github.io/pkg/resume/svc"
)

var (
	searchLimit      int
	searchPage       int
	searchAll        int
	searchSalary     int
	searchCategories []string
	searchLevels     []string
	searchTypes      []string
	searchPostedBy   []string
	searchFlexible   []int
	searchSort       string
)

// MCFSearchCmd 查詢職缺。結果不含 JD 全文，需要全文請接 mcf detail。
var MCFSearchCmd = &cobra.Command{
	Use:   "search [keyword]",
	Short: "Search job postings by keyword and filters",
	Long: "search 對應 POST /v2/search。\n" +
		"回傳的職缺不含 description —— 那是 API 的行為，不是本命令的裁切，要 JD 全文請接 mcf detail 或改用 mcf fetch。\n" +
		"--category / --level / --employment / --posted-by 都是伺服器端 enum，值必須與網站顯示名稱完全一致，錯字會換來 400 而非空結果。\n" +
		"--sort 只接受 new_posting_date 與 min_monthly_salary，留白代表以相關性排序。\n" +
		"API 另有 schemes 與 jobStatuses 兩個欄位，前者無作用、後者會讓其餘篩選全部失效，因此本命令刻意不提供。",
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(cmd.Context(), commandTimeout())
		defer cancel()

		request := svc.SearchRequest{
			Salary:                   searchSalary,
			Categories:               searchCategories,
			PositionLevels:           searchLevels,
			EmploymentTypes:          searchTypes,
			PostingCompany:           searchPostedBy,
			FlexibleWorkArrangements: searchFlexible,
		}
		if len(args) == 1 {
			request.Search = args[0]
		}
		if searchSort != "" {
			request.SortBy = []string{searchSort}
		}
		client := newMCFClient()
		if searchAll > 0 {
			jobs, err := client.SearchAll(ctx, request, searchAll)
			if err != nil {
				return err
			}
			return renderJobs(jobs)
		}

		result, err := client.Search(ctx, request, svc.Page{Limit: searchLimit, Page: searchPage})
		if err != nil {
			return err
		}
		if jsonFlag {
			return writeJSON(os.Stdout, result)
		}
		writeJobTable(os.Stdout, result.Results)
		cmd.Printf("\n%d matched (%d without filters)\n", result.Total, result.CountWithoutFilters)
		return nil
	},
}

// renderJobs 是 --all 路徑的輸出，沒有 total 可報，只印結果本身。
func renderJobs(jobs []svc.Job) error {
	if jsonFlag {
		return writeJSON(os.Stdout, jobs)
	}
	writeJobTable(os.Stdout, jobs)
	return nil
}

func init() {
	flags := MCFSearchCmd.Flags()
	flags.IntVar(&searchLimit, "limit", 20, "results per page (server maximum 100)")
	flags.IntVar(&searchPage, "page", 0, "zero-based page index")
	flags.IntVar(&searchAll, "all", 0, "page through up to N results, ignoring --limit/--page")
	flags.IntVar(&searchSalary, "salary", 0, "minimum monthly salary")
	flags.StringSliceVar(&searchCategories, "category", nil, "job category, e.g. 'Information Technology'")
	flags.StringSliceVar(&searchLevels, "level", nil, "position level, e.g. Manager, Senior Executive")
	flags.StringSliceVar(&searchTypes, "employment", nil, "employment type, e.g. 'Full Time', Permanent")
	flags.StringSliceVar(&searchPostedBy, "posted-by", nil, "'Direct' (employer) or 'Third Party' (recruiter)")
	flags.IntSliceVar(&searchFlexible, "flexible", nil, "flexible work arrangement id: 1 Flexi-Hours, 2 Telecommuting, 3 Employees Choice of Days Off, 4 Staggered Time, 5 Compressed Work Schedule, 6 Creative Scheduling")
	flags.StringVar(&searchSort, "sort", "", "new_posting_date | min_monthly_salary (default: relevancy)")
	flags.BoolVar(&jsonFlag, "json", false, "emit raw JSON instead of a table")

	MCFCmd.AddCommand(MCFSearchCmd)
}
