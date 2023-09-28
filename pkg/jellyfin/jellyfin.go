package jellyfin

import (
	"context"
	"fmt"
	"path"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//go:generate oapi-codegen --package=jellyfin -generate=client,types -o ./jellyfin.gen.go https://jellyfin.kielbasa.link/api-docs/openapi.json

func Jellyfin(cmd *cobra.Command, args []string) error {
  tokenString := fmt.Sprintf("MediaBrowser Token=\"%s\"", viper.GetString("jellyfin_api_key"))
	apiKey, err := securityprovider.NewSecurityProviderApiKey("header", "Authorization", tokenString)
	c, err := NewClientWithResponses(viper.GetString("jellyfin_url"), WithRequestEditorFn(apiKey.Intercept))
	if err != nil {
    return err
	}
	params := GetSessionsParams{}
	resp, err := c.GetSessionsWithResponse(context.Background(), &params)
	if err != nil {
    return err
	}
  if resp.StatusCode() != 200 {
    err := fmt.Errorf("Failed getting Jellyfin sessions: %s\n", resp.Status())
    return err
  }
	sessions := *resp.JSON200
	for _, session := range sessions {
		if session.NowPlayingItem != nil {
			item := *session.NowPlayingItem
			file := path.Base(*item.Path)
			var status string
			if *session.PlayState.IsPaused {
				status = "paused"
			} else {
				status = "playing"
			}
      progress := float64(*session.PlayState.PositionTicks) / float64(*item.RunTimeTicks) * 100
			cmd.Printf("[%v] [%v] [Status: %v (%0.0f%%)]\n", *session.UserName, file, status, progress)
		}
	}

  return nil
}
