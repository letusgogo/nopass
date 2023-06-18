/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// verboseCmd represents the verbose command
var verboseCmd = &cobra.Command{
	Use:   "gen",
	Short: "generate password",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gen called")
	},
}

func init() {
	rootCmd.AddCommand(verboseCmd)
}
