package cmd

import (
	"context"
	"os"

	"github.com/spf13/cobra"

	"github.com/BizShuk/bizshuk.github.io/pkg/resume/svc"
)

var (
	companyLimit      int
	companyPage       int
	companyUEN        string
	companyResponsive bool
)

// MCFCompanyCmd 查詢雇主，用於把 JD 上的公司名稱解析成 UEN 後再抓該公司的職缺。
var MCFCompanyCmd = &cobra.Command{
	Use:   "company [name]",
	Short: "Look up employers by name or UEN",
	Long: "company 對應 GET /v2/companies 與 /v2/companies/{uen}。\n" +
		"回應的 _links 另指向 jobs / addresses / schemes 三條子資源，但那三條需要授權，未帶憑證會回 401。\n" +
		"公開路徑上沒有「列出某雇主的所有職缺」，只能用 mcf search 以雇主名稱當關鍵字近似取得。",
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(cmd.Context(), commandTimeout())
		defer cancel()

		client := newMCFClient()
		page := svc.Page{Limit: companyLimit, Page: companyPage}

		if companyUEN != "" {
			company, err := client.Company(ctx, companyUEN)
			if err != nil {
				return err
			}
			if jsonFlag {
				return writeJSON(os.Stdout, company)
			}
			writeCompanyTable(os.Stdout, []svc.Company{*company})
			return nil
		}

		query := svc.CompanyQuery{
			OrderBy:            svc.COMPANY_ORDER_BY_NAME,
			OrderDirection:     "asc",
			ResponsiveEmployer: companyResponsive,
			Page:               page,
		}
		if len(args) == 1 {
			query.Name = args[0]
		}

		list, err := client.Companies(ctx, query)
		if err != nil {
			return err
		}
		if jsonFlag {
			return writeJSON(os.Stdout, list)
		}
		writeCompanyTable(os.Stdout, list.Results)
		cmd.Printf("\n%d employers matched\n", list.Total)
		return nil
	},
}

func init() {
	flags := MCFCompanyCmd.Flags()
	flags.IntVar(&companyLimit, "limit", 20, "results per page (server maximum 100)")
	flags.IntVar(&companyPage, "page", 0, "zero-based page index")
	flags.StringVar(&companyUEN, "uen", "", "look up one employer by its Unique Entity Number")
	flags.BoolVar(&companyResponsive, "responsive", false, "only employers flagged as responsive to applicants")
	flags.BoolVar(&jsonFlag, "json", false, "emit raw JSON instead of a table")

	MCFCmd.AddCommand(MCFCompanyCmd)
}
