package cmd

import (
	"context"
	"os"

	"github.com/spf13/cobra"

	"github.com/BizShuk/bizshuk.github.io/pkg/resume/svc"
)

var (
	jobsLimit  int
	jobsPage   int
	jobsSalary int
)

// MCFJobsCmd 列出全站職缺。與 search 不同，這條路徑帶 description 但不吃關鍵字。
var MCFJobsCmd = &cobra.Command{
	Use:   "jobs",
	Short: "List all postings site-wide, newest first",
	Long: "jobs 對應 GET /v2/jobs。\n" +
		"它不吃關鍵字但直接帶回 description，適合全量掃描；要篩選請用 mcf search。",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(cmd.Context(), commandTimeout())
		defer cancel()

		list, err := newMCFClient().Jobs(ctx, svc.Page{Limit: jobsLimit, Page: jobsPage}, jobsSalary)
		if err != nil {
			return err
		}
		if jsonFlag {
			return writeJSON(os.Stdout, list)
		}
		writeJobTable(os.Stdout, list.Results)
		cmd.Printf("\n%d postings site-wide\n", list.Total)
		return nil
	},
}

func init() {
	flags := MCFJobsCmd.Flags()
	flags.IntVar(&jobsLimit, "limit", 20, "results per page (server maximum 100)")
	flags.IntVar(&jobsPage, "page", 0, "zero-based page index")
	flags.IntVar(&jobsSalary, "salary", 0, "minimum monthly salary")
	flags.BoolVar(&jsonFlag, "json", false, "emit raw JSON instead of a table")

	MCFCmd.AddCommand(MCFJobsCmd)
}
