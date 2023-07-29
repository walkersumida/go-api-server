package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/walkersumida/go-api-server/cmd/db/migrate"
)

func main() {
	rootCmd := &cobra.Command{Use: "db"}
	rootCmd.AddCommand(cmdCreate)
	rootCmd.AddCommand(cmdDrop)
	rootCmd.AddCommand(cmdMigrate)
	cmdMigrate.AddCommand(migrate.CmdMigrateUp)
	cmdMigrate.AddCommand(migrate.CmdMigrateDown)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
