package main

import "github.com/spf13/cobra"

var cmdMigration = &cobra.Command{
	Use:   "migration",
	Short: "migration files",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
