/*
Copyright © 2023 Helloworld helloworldyong9@gmail.com
*/

package cmd

import (
	"github.com/letusgogo/nopass/config"
	"github.com/letusgogo/nopass/log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	defaultConfig *config.Config
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nopass",
	Short: "A password generator based on user-defined fields (nopass)",

	Run: func(cmd *cobra.Command, args []string) {

	},
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
	cobra.OnInitialize(initLog)
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringP("config", "c", "", "config file (default is $HOME/.nopass.yaml)")
	rootCmd.PersistentFlags().StringP("log", "l", "info", "log level (default is info)")
}

func initLog() {
	logLevel, err := rootCmd.Flags().GetString("log")
	if err != nil {
		panic(err)
	}
	log.SetLevel(logLevel)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	cfgFile, err := rootCmd.Flags().GetString("config")
	if err != nil {
		panic(err)
	}
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".nopass" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".nopass")
	}
	// read in environment variables that match
	viper.AutomaticEnv()

	log.Debug("using config file: ", viper.ConfigFileUsed())

	// If a config file is found, read it in.
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read %s failed: %v", viper.ConfigFileUsed(), err)
	}

	defaultConfig, err = config.LoadConfig()
	if err != nil {
		log.Fatalf("load %s failed: %v", viper.ConfigFileUsed(), err)
	}
	log.DrawPhase("config loaded", log.DebugLevel, func() {
		log.Debug(defaultConfig)
	})
}
