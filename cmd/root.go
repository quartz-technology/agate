package cmd

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
//
//nolint:exhaustruct,gochecknoglobals
var rootCmd = &cobra.Command{
	Use:   "agate",
	Short: "An Ethereum Relay Data Indexer",
	Long: `
Agate is an application which aggregates data from relays on the ethereum network and stores it
in a database in an attempt to enlighten a bit more the dark forest.
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

//nolint:gochecknoglobals
var configurationFilePath string

func init() {
	cobra.OnInitialize(initConfig)

	// Registering specific root flags.
	rootCmd.PersistentFlags().StringVar(&configurationFilePath, "config", "", "path to the configuration file")
}

func initConfig() {
	if configurationFilePath != "" {
		// Use config file from the flag.
		viper.SetConfigFile(configurationFilePath)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".agate-config" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".agate-config")
	}

	viper.SetEnvPrefix("AGATE")
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Debug().
			Str("file", viper.ConfigFileUsed()).
			Msg("using configuration file")
	}
}
