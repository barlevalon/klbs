package plex

import (
	"github.com/spf13/cobra"
)

var PlexCmd = &cobra.Command{
	Use:   "plex",
	Short: "plex commands",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
}
