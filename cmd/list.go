/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"fmt"
	"github.com/letusgogo/nopass/log"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var (
	ruleName string
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list the elements of select rule",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debugf("list called, args: %v\n", args)

		if defaultConfig == nil {
			fmt.Println("can not load config file, please check it in ~/.nopass.yaml")
			return
		}
		var (
			rulesInYaml []byte
			err         error
		)
		// display all the rules in yaml
		if ruleName == "" {
			rulesInYaml, err = yaml.Marshal(defaultConfig.Rules)
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			if rule, ok := defaultConfig.Rules[ruleName]; ok {
				rulesInYaml, err = yaml.Marshal(rule)
				if err != nil {
					fmt.Println(err)
					return
				}
			} else {
				fmt.Printf("can not find rule: %s\n", ruleName)
				return
			}
		}
		fmt.Println(string(rulesInYaml))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&ruleName, "rule", "r", "", `list rule info, eg: -r=difficult`)
}
