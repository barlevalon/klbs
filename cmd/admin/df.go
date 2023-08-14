package admin

import (
	"github.com/barlevalon/klbs/lib/remote"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var dfCmd = &cobra.Command{
	Use:   "df",
	Short: "",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
    remote.Exec(viper.GetString("host"), "df -h /")
	},
}

func init() {
	AdminCmd.AddCommand(dfCmd)
}
