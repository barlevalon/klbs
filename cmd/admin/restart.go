package admin

import (
	"context"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

  "github.com/barlevalon/klbs/lib/remote"
)

var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restarts main services in correct order",
	Long: `Restarts rclone-mount, plex, nzbget, rtorrent-flood, sonarr, radarr. In order.`,
  Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
    docker := remote.DockerSSHClient{Host: viper.GetString("host")}.Client()
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    allTargets := []string{"rclone-mount", "plex", "nzbget", "rtorrent-flood", "sonarr", "radarr"}
    targets := []string{}
    if len(args) == 1 && args[0] == "all" {
      targets = append(targets, allTargets...)
    } else {
      targets = append(targets, args...)
    }

    for _, target := range targets {
      err := docker.ContainerRestart(ctx, target, container.StopOptions{})
      if err != nil {
        cmd.PrintErrf("Error restarting %s: %s\n", target, err)
      } else {
        cmd.Printf("Restarted %s\n", target)
      }
    }
	},
}

func init() {
  AdminCmd.AddCommand(restartCmd)
}
