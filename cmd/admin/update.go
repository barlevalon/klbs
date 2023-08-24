package admin

import (
	"errors"
	"fmt"
	"slices"

	"github.com/barlevalon/klbs/pkg/tautulli"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update components",
	Long:  ``,
	Args:  validateArgs,
	Run: func(cmd *cobra.Command, args []string) {
		activity, err := tautulli.GetActivity()
		if err != nil {
			cmd.PrintErrf("Error getting activity: %s\n", err)
		}
		if activity.Response.Data.StreamCount != "0" {
			prompt := promptui.Prompt{
				Label:     "There are active streams. Are you sure you want to update?",
				IsConfirm: true,
			}
			result, _ := prompt.Run()
			if result != "y" {
				return
			}
		}
		cmd.Printf("Updating...")
    // If args is "all", set args to all targets
    // If args contains "os", run apt update && apt upgrade
    // If args contains "rclone-mount", set restart required
    // Run watchtower update for each image
    // If reboot required, set restart required, prompt to reboot
    // If restart required, restart
	},
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("update requires at least one arg")
	}

	if slices.Contains(args, "all") && len(args) > 1 {
		return errors.New("cannot specify 'all' with other targets")
	}
 
  validTargets := []string{"all", "rclone-mount", "plex", "nzbget", "rtorrent-flood", "sonarr", "radarr", "tautulli", "prowlarr", "overseerr", "os"}
	for _, arg := range args {
		if !slices.Contains(validTargets, arg) {
			return errors.New(fmt.Sprintf("invalid target: %s", arg))
		}
	}
	return nil
}

func init() {
	AdminCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
