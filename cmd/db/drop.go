package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/walkersumida/go-api-server/database"
)

var cmdDrop = &cobra.Command{
	Use:   "drop",
	Short: "Drop database",
	Run: func(cmd *cobra.Command, args []string) {
		dropDB()
	},
}

func dropDB() {
	configs := database.Configs{Name: "postgres"}
	db, err := database.NewClient(configs)
	if err != nil {
		panic(err)
	}

	name := os.Getenv("DB_NAME")

	err = db.Drop(name)
	if err != nil {
		panic(err)
	}

	println("Database dropped")
}
