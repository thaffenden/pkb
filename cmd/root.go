// Package cmd contains the different CLI commands for interactions in pkb.
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thaffenden/pkb/internal/config"
	"github.com/thaffenden/pkb/internal/flags"
)

// Version is the CLI version set via linker flags at build time.
var Version string

var rootCmd = &cobra.Command{
	RunE: func(ccmd *cobra.Command, args []string) error {
		return nil
	},
	Short:   "manage notes in markdown files",
	Use:     "pkb",
	Version: Version,
}

// Execute executes the root command.
func Execute(conf config.Config) error {
	ctx := context.WithValue(context.Background(), config.ContextKey, conf)
	return rootCmd.ExecuteContext(ctx)
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(CreateNew())
	rootCmd.AddCommand(CreateEdit())
	rootCmd.AddCommand(CreateCopy())
	rootCmd.PersistentFlags().StringVar(&flags.ConfigFile, "config", "", "config file if not held at default location")
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
}

func initConfig() {
	if flags.ConfigFile == "" {
		viper.SetConfigName("config")
		viper.SetConfigType("json")
		viper.AddConfigPath("$XDG_CONFIG_DIR/pkb/")
		viper.AddConfigPath("$HOME/.config/pkb/")
	} else {
		viper.SetConfigFile(flags.ConfigFile)
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("error trying to read config file: %s", err)
		os.Exit(1)
	}
}
