// Package cmd 是 resume 的命令集合，一個檔案對應一個命令。
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/bizshuk/gosdk/config"
	gohttp "github.com/bizshuk/gosdk/http"
	_ "github.com/bizshuk/gosdk/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/BizShuk/bizshuk.github.io/pkg/resume/svc"
)

// APP_NAME 決定設定與資料目錄：~/.config/resume-jd 下的 data/ 與 logs/。
const APP_NAME = "resume-jd"

// 扁平 viper key。全部可由同名大寫環境變數覆寫。
const (
	KEY_MCF_BASE_URL     = "MCF_BASE_URL"
	KEY_MCF_USER_AGENT   = "MCF_USER_AGENT"
	KEY_MCF_TIMEOUT      = "MCF_TIMEOUT"
	KEY_MCF_MAX_ATTEMPTS = "MCF_MAX_ATTEMPTS"
)

// RootCmd 是 CLI 進入點，各資料來源掛在它底下作為子命令。
var RootCmd = &cobra.Command{
	Use:               "resume",
	Short:             "Collect job descriptions for the resume job library",
	Long:              "resume 蒐集外部職缺資料，供 pkg/resume/jd 的職缺庫與匹配度評估使用。",
	SilenceUsage:      true,
	SilenceErrors:     true,
	PersistentPreRunE: loadConfig,
}

// Execute 是 main 的唯一進入點。
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

// loadConfig 交給 gosdk 決定設定來源：工作目錄、./conf 與 ~/.config/resume-jd，
// 每層都吃 .env 與 .env.local，process env 優先權最高。
func loadConfig(_ *cobra.Command, _ []string) error {
	config.Default(config.WithAppName(APP_NAME))
	return nil
}

// newMCFClient 是 MyCareersFuture client 的唯一建構入口，設定一律走扁平大寫 key。
func newMCFClient() *svc.Client {
	opts := []svc.Option{
		svc.WithBaseURL(viper.GetString(KEY_MCF_BASE_URL)),
		svc.WithUserAgent(viper.GetString(KEY_MCF_USER_AGENT)),
	}
	if timeout := viper.GetDuration(KEY_MCF_TIMEOUT); timeout > 0 {
		opts = append(opts, svc.WithTimeout(timeout))
	}
	if attempts := viper.GetInt(KEY_MCF_MAX_ATTEMPTS); attempts > 0 {
		opts = append(opts, svc.WithRetryPolicy(gohttp.RetryPolicy{
			MaxAttempts: attempts,
			BaseDelay:   gohttp.DEFAULT_BASE_DELAY,
			MaxDelay:    gohttp.DEFAULT_MAX_DELAY,
		}))
	}
	return svc.New(opts...)
}

// commandTimeout 給每個命令一個上限，避免全量抓取無限期掛住。
func commandTimeout() time.Duration {
	if timeout := viper.GetDuration("COMMAND_TIMEOUT"); timeout > 0 {
		return timeout
	}
	return 10 * time.Minute
}
