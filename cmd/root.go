package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "repoview",
	Short: "A tool to view local repositories",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $XDG_CONFIG_HOME/repoview/config.yaml)")
	rootCmd.PersistentFlags().String("storage-path", "", "path to storage file")
	viper.BindPFlag("storage.path", rootCmd.PersistentFlags().Lookup("storage"))

	// Set default values
	defaultStoragePath := filepath.Join(xdg.DataHome, "repoview", "storage.json")
	viper.SetDefault("storage.path", defaultStoragePath)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Use XDG config directory
		configDir := filepath.Join(xdg.ConfigHome, "repoview")
		if err := os.MkdirAll(configDir, 0755); err != nil {
			cobra.CheckErr(err)
		}

		// Search config in XDG config directory
		viper.AddConfigPath(configDir)
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
