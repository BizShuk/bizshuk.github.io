package cmd

import (
	"context"
	"os"

	"github.com/spf13/cobra"
)

var detailQuestions bool

// MCFDetailCmd 取回單一職缺的完整內容，這是唯一拿得到 JD 全文的路徑。
var MCFDetailCmd = &cobra.Command{
	Use:   "detail <uuid>",
	Short: "Fetch one job posting in full, including the description",
	Long: "detail 對應 GET /v2/jobs/{uuid}。\n" +
		"uuid 來自 mcf search 的第一欄，也可從職缺頁網址結尾取得。",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(cmd.Context(), commandTimeout())
		defer cancel()

		client := newMCFClient()
		job, err := client.Job(ctx, args[0])
		if err != nil {
			return err
		}

		if detailQuestions {
			questions, err := client.ScreeningQuestions(ctx, args[0])
			if err != nil {
				return err
			}
			job.ScreeningQuestions = questions
		}

		if jsonFlag {
			return writeJSON(os.Stdout, job)
		}
		writeJobDetail(os.Stdout, job)
		return nil
	},
}

func init() {
	flags := MCFDetailCmd.Flags()
	flags.BoolVar(&detailQuestions, "questions", false, "also fetch the application screening questions")
	flags.BoolVar(&jsonFlag, "json", false, "emit raw JSON instead of a readable summary")

	MCFCmd.AddCommand(MCFDetailCmd)
}
