/*
Copyright © 2023 Helloworld helloworldyong9@gmail.com
*/

package cmd

import (
	"bytes"
	"fmt"
	"github.com/letusgogo/nopass/config"
	"github.com/letusgogo/nopass/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"os"
)

var (
	defaultConfig *config.Config
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nopass",
	Short: "A password generator based on user-defined fields (nopass)",

	Run: func(cmd *cobra.Command, args []string) {
		getPass()
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

	rootCmd.PersistentFlags().StringP("config", "c", "nopass.yaml", "config file (default is $HOME/nopass.yaml)")
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
	viper.SetConfigFile(cfgFile)

	// check if config file exists, if not, create one
	if _, err = os.Stat(cfgFile); os.IsNotExist(err) {
		hintCreateConfig()
	} else {
		// If a config file is found, read it in.
		err = viper.ReadInConfig()
		if err != nil {
			log.Fatalf("read %s failed: %v", viper.ConfigFileUsed(), err)
		}
	}
	log.Debug("using config file: ", viper.ConfigFileUsed())

	defaultConfig, err = config.LoadConfig()
	if err != nil {
		log.Fatalf("load %s failed: %v\n", viper.ConfigFileUsed(), err)
	}
	log.DrawPhase("config loaded", log.DebugLevel, func() {
		out, err := yaml.Marshal(defaultConfig)
		if err != nil {
			log.Fatal(err)
		}
		log.Debug(string(out))
	})
}

func hintCreateConfig() {
	log.Hint("can not find config file, create a new one.")
	log.Hint("1: for English, 2: 中文配置, 3: exit")
	var (
		input string
		lang  string
	)
	_, _ = fmt.Scanln(&input)
	if input == "1" {
		lang = "en"
	} else if input == "2" {
		lang = "zh"
	} else {
		log.Fatal("can not find config file")
	}

	content, err := config.GetDefaultConfig(lang)
	if err != nil {
		log.Fatal(err)
	}

	err = viper.ReadConfig(bytes.NewBufferString(content))
	if err != nil {
		log.Fatalf("read default config failed: %v", err)
	}
	err = viper.WriteConfig()
	if err != nil {
		log.Fatalf("write default config failed: %v", err)
	}
}
