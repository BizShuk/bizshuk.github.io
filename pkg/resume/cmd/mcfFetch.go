package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"path/filepath"
	"time"

	"github.com/bizshuk/gosdk/config"
	"github.com/bizshuk/gosdk/file"
	"github.com/spf13/cobra"

	"github.com/BizShuk/bizshuk.github.io/pkg/resume/svc"
)

var (
	fetchMax     int
	fetchSalary  int
	fetchLevels  []string
	fetchOutDir  string
	fetchPace    time.Duration
	fetchRefresh bool
)

// MCFFetchCmd 把「搜尋拿 uuid，再逐筆補 JD 全文」這條必經的兩段式流程包成一個命令。
var MCFFetchCmd = &cobra.Command{
	Use:   "fetch <keyword>",
	Short: "Search, then fetch each posting in full and write it to disk",
	Long: "fetch 先 POST /v2/search 收集 uuid，再逐筆 GET /v2/jobs/{uuid} 補回 description，\n" +
		"最後以 <uuid>.json 落檔。兩段式是 API 的形狀決定的：search 的結果永遠不含 JD 全文。\n" +
		"每筆之間預設間隔 --pace，避免觸發未公開的流量限制。",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(cmd.Context(), commandTimeout())
		defer cancel()

		store, err := file.NewStore[svc.Job](fetchOutputDir())
		if err != nil {
			return err
		}

		client := newMCFClient()
		summaries, err := client.SearchAll(ctx, svc.SearchRequest{
			Search:         args[0],
			Salary:         fetchSalary,
			PositionLevels: fetchLevels,
			SortBy:         []string{svc.SORT_NEW_POSTING_DATE},
		}, fetchMax)
		if err != nil {
			return err
		}

		cmd.Printf("%d postings matched, writing to %s\n", len(summaries), store.Dir())

		var failed int
		for index, summary := range summaries {
			if !fetchRefresh && store.Exists(summary.UUID) {
				slog.Debug("skip cached posting", "uuid", summary.UUID, "title", summary.Title)
				continue
			}
			if index > 0 && fetchPace > 0 {
				if err := pause(ctx, fetchPace); err != nil {
					return err
				}
			}

			job, err := client.Job(ctx, summary.UUID)
			if err != nil {
				if ctx.Err() != nil {
					return err
				}
				failed++
				slog.Error("fetch posting failed", "uuid", summary.UUID, "error", err)
				continue
			}
			if err := store.Write(job.UUID, *job); err != nil {
				return err
			}
			cmd.Printf("  %s  %s\n", job.UUID, truncate(job.Title, 60))
		}

		if failed > 0 {
			return fmt.Errorf("%d of %d postings could not be fetched", failed, len(summaries))
		}
		return nil
	},
}

// fetchOutputDir 決定落檔位置：未指定時使用 gosdk 的 app data 目錄。
func fetchOutputDir() string {
	if fetchOutDir != "" {
		return fetchOutDir
	}
	return filepath.Join(config.GetAppDataDir(), "mcf", "jobs")
}

// pause 是可被 context 中止的等待，讓 Ctrl-C 不必等滿一個間隔才生效。
func pause(ctx context.Context, d time.Duration) error {
	timer := time.NewTimer(d)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

func init() {
	flags := MCFFetchCmd.Flags()
	flags.IntVar(&fetchMax, "max", 50, "stop after this many postings (0 = every match)")
	flags.IntVar(&fetchSalary, "salary", 0, "minimum monthly salary")
	flags.StringSliceVar(&fetchLevels, "level", nil, "position level, e.g. Manager, Senior Executive")
	flags.StringVar(&fetchOutDir, "out", "", "output directory (default: <app data dir>/mcf/jobs)")
	flags.DurationVar(&fetchPace, "pace", 300*time.Millisecond, "delay between detail requests")
	flags.BoolVar(&fetchRefresh, "refresh", false, "re-fetch postings already on disk")

	MCFCmd.AddCommand(MCFFetchCmd)
}
