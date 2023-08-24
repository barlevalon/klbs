package plex

import (
	"fmt"

	"github.com/barlevalon/klbs/pkg/tautulli"
	"github.com/spf13/cobra"
)


var sessionsCmd = &cobra.Command{
	Use:   "sessions",
	Short: "List active sessions",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
    getActivityResponse, err := tautulli.GetActivity()
    if err != nil {
      cmd.PrintErrf("Error getting activity: %s\n", err)
    }
		cmd.Printf("Streams: %s\n", getActivityResponse.Response.Data.StreamCount)
		for _, session := range getActivityResponse.Response.Data.Sessions {
			var title string
			if session.MediaType == "episode" {
				title = fmt.Sprintf("%s (%s) S%sE%s", session.GrandparentTitle, session.GrandparentYear, session.ParentMediaIndex, session.MediaIndex)
			} else {
				title = fmt.Sprintf("%s (%s)", session.Title, session.Year)
			}
			cmd.Printf("[%s] [%s] [Quality: %s] [Status: %s (%s%%)]\n", session.User, title, session.QualityProfile, session.State, session.ProgressPercent)
		}
	},
}

func init() {
	PlexCmd.AddCommand(sessionsCmd)
}
