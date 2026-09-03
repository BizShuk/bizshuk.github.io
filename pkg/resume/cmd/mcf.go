package cmd

import "github.com/spf13/cobra"

// MCFCmd 收攏所有 MyCareersFuture 端點。它本身不做事，只作為子命令的掛載點。
var MCFCmd = &cobra.Command{
	Use:   "mcf",
	Short: "Query the MyCareersFuture public API",
	Long: "mcf 讀取 api.mycareersfuture.gov.sg。\n" +
		"該 API 不需要 API key，但屬於網站前端使用的非官方介面，端點與欄位可能無預警改版。",
}

func init() {
	RootCmd.AddCommand(MCFCmd)
}
