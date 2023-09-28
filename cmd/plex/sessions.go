package plex

import (
	"fmt"

	"github.com/barlevalon/klbs/pkg/jellyfin"
	"github.com/barlevalon/klbs/pkg/tautulli"
	"github.com/spf13/cobra"
)


var sessionsCmd = &cobra.Command{
	Use:   "sessions",
	Short: "List active sessions",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
    {
      err := jellyfin.Jellyfin(cmd, args)
      if err != nil {
        return err
      }
    }
    getActivityResponse, err := tautulli.GetActivity()
    if err != nil {
      return fmt.Errorf("Failed getting activity: %v\n", err)
    }
		cmd.Printf("Plex streams: %s\n", getActivityResponse.Response.Data.StreamCount)
		for _, session := range getActivityResponse.Response.Data.Sessions {
			var title string
			if session.MediaType == "episode" {
				title = fmt.Sprintf("%s (%s) S%sE%s", session.GrandparentTitle, session.GrandparentYear, session.ParentMediaIndex, session.MediaIndex)
			} else {
				title = fmt.Sprintf("%s (%s)", session.Title, session.Year)
			}
			cmd.Printf("[%s] [%s] [Quality: %s] [Status: %s (%s%%)]\n", session.User, title, session.QualityProfile, session.State, session.ProgressPercent)
		}

    return nil
	},
}

func init() {
	PlexCmd.AddCommand(sessionsCmd)
}
