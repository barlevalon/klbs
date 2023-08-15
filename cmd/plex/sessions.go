package plex

import (
	"fmt"

	"github.com/barlevalon/klbs/lib/tautulli"
	"github.com/imroc/req/v3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type ErrorMessage struct {
	Response struct {
		Data    struct{} `json:"data"`
		Result  string   `json:"result"`
		Message string   `json:"message"`
	} `json:"response"`
}

var sessionsCmd = &cobra.Command{
	Use:   "sessions",
	Short: "List active sessions",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var getActivityResponse tautulli.GetActivityResponse
		var errMsg ErrorMessage
		client := req.C().
			SetBaseURL(viper.GetString("tautulli_api_url")).
			SetCommonQueryParams(map[string]string{"apikey": viper.GetString("tautulli_api_key")})
		resp, err := client.R().
			SetSuccessResult(&getActivityResponse).
      SetErrorResult(&errMsg).
			Get("?cmd=get_activity")
		if err != nil {
			cmd.PrintErrln("error:", err)
			cmd.PrintErrln("raw content:")
			cmd.PrintErrln(resp.Dump())
			return
		}
		if resp.IsErrorState() {
      cmd.PrintErrf("Tautulli error: %s\n", errMsg.Response.Message)
			return
		}

		if resp.IsSuccessState() {
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
			return
		}

		cmd.PrintErrln("unknown status", resp.Status)
		cmd.PrintErrln("raw content:")
		cmd.PrintErrln(resp.Dump())
	},
}

func init() {
	PlexCmd.AddCommand(sessionsCmd)
}
