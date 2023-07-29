package main

import "github.com/spf13/cobra"

var cmdMigrate = &cobra.Command{
	Use:   "migrate",
	Short: "migrate database",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
