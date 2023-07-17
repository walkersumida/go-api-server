package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/walkersumida/go-api-server/database"
)

var cmdCreate = &cobra.Command{
	Use:   "create",
	Short: "Create database",
	Run: func(cmd *cobra.Command, args []string) {
		createDB()
	},
}

func createDB() {
	configs := database.Configs{Name: "postgres"}
	db, err := database.NewClient(configs)
	if err != nil {
		panic(err)
	}

	name := os.Getenv("DB_NAME")

	err = db.Create(name)
	if err != nil {
		panic(err)
	}

	println("Database created")
}
