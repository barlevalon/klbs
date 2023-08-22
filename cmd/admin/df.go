package admin

import (
	"github.com/barlevalon/klbs/pkg/remote"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var dfCmd = &cobra.Command{
	Use:   "df",
	Short: "",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
    output, err := remote.Exec(viper.GetString("host"), "df -h /")
    if err != nil {
      cmd.PrintErrf("Error running df: %s\n", err)
    }
    cmd.Printf("%s", output)
	},
}

func init() {
	AdminCmd.AddCommand(dfCmd)
}
