package migrate

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
	"github.com/walkersumida/go-api-server/database"
	"github.com/walkersumida/go-api-server/internal/pkg/path"
)

var CmdMigrateDown = &cobra.Command{
	Use:   "down",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		down()
	},
}

func down() {
	url := database.BuildDBURL(database.Configs{})
	m, err := migrate.New(
		"file://"+path.Migration(),
		url)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Down(); err != nil {
		log.Fatal(err)
	}

	println("Database migrated")
}
