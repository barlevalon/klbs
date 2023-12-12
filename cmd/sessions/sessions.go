package sessions

import (
	"github.com/barlevalon/klbs/pkg/jellyfin"
	"github.com/spf13/cobra"
)


var SessionsCmd = &cobra.Command{
	Use:   "sessions",
	Short: "List active sessions",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
    {
      err := jellyfin.PrintJellyfinSessions(cmd, args)
      if err != nil {
        return err
      }
    }
    return nil
	},
}

func init() {
}
