/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"fmt"
	"github.com/letusgogo/nopass/algo"
	"github.com/letusgogo/nopass/gen"
	"github.com/letusgogo/nopass/log"
	"github.com/letusgogo/nopass/rule"
	"github.com/spf13/cobra"
)

var (
	genRuleName string
	genAlgo     string
)

// generateCmd represents the verbose command
var generateCmd = &cobra.Command{
	Use:   "gen",
	Short: "generate password",
	Run: func(cmd *cobra.Command, _ []string) {
		log.Debug("gen called")

		algoConfig := defaultConfig.Algo[genAlgo]
		if algoConfig == nil {
			log.Errorf("can not find algorithmFromConfig: %s\n", genAlgo)
			return
		}
		algorithmFromConfig, err := algo.NewAlgorithmFromConfig(algoConfig)
		if err != nil {
			log.Errorf("can not create algorithmFromConfig: %s\n", genAlgo)
			return
		}

		ruleFromConfig, err := rule.NewRuleFromConfig(genRuleName, defaultConfig)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = ruleFromConfig.FullElements()
		if err != nil {
			log.Error(err)
		}
		ruleFromConfig.Display()

		password := gen.GeneratePassword(ruleFromConfig, algorithmFromConfig)
		log.Hint(password)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&genRuleName, "rule", "r", "default", `select rule to generate password, eg: -r=simple`)
	generateCmd.Flags().StringVarP(&genAlgo, "algo", "a", "sha256", `select algo to generate password, eg: -a=sha256`)
}
