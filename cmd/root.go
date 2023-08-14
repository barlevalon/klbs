package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

  "github.com/barlevalon/klbs/cmd/admin"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "klbs",
	Short: "Control kielbasa. Make that sausage.",
	Long: ``,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/.klbs.yaml)")
  rootCmd.PersistentFlags().StringP("host", "H", viper.GetString("host"), "Host to connect to")

  rootCmd.AddCommand(admin.AdminCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home + "/.config/klbs/")
		viper.SetConfigType("toml")
		viper.SetConfigName("klbs")
	}
  viper.BindPFlag("host", rootCmd.PersistentFlags().Lookup("host"))
  viper.SetEnvPrefix("klbs")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
  if viper.GetString("host") == "" {
    cobra.CheckErr("Host not set")
  }
}
