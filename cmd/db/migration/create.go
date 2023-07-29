package migration

import (
	"context"
	"log"
	"os"

	"ariga.io/atlas/sql/sqltool"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"github.com/walkersumida/go-api-server/database"
	"github.com/walkersumida/go-api-server/ent/migrate"
	"github.com/walkersumida/go-api-server/internal/pkg/path"
)

var CmdMigrationCreate = &cobra.Command{
	Use:   "create",
	Short: "Create migrations",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		create(args[0])
	},
}

func create(fileName string) {
	ctx := context.Background()

	dir, err := sqltool.NewGolangMigrateDir(path.Migration())
	if err != nil {
		log.Fatalf("failed creating atlas migration directory: %v", err)
	}

	// Migrate diff options.
	opts := []schema.MigrateOption{
		schema.WithDir(dir),                         // provide migration directory
		schema.WithMigrationMode(schema.ModeReplay), // provide migration mode
		schema.WithDialect(dialect.Postgres),        // Ent dialect to use
	}

	configs := database.Configs{Name: "postgres"}
	db, err := database.NewClient(configs)
	if err != nil {
		panic(err)
	}

	dbname := os.Getenv("DB_NAME") + "_for_migration"

	err = db.Drop(dbname)
	if err != nil {
		panic(err)
	}

	err = db.Create(dbname)
	if err != nil {
		panic(err)
	}

	url := database.BuildDBURL(database.Configs{Name: dbname})
	err = migrate.NamedDiff(ctx, url, fileName, opts...)
	if err != nil {
		log.Fatalf("failed generating migration file: %v", err)
	}

	err = db.Drop(dbname)
	if err != nil {
		panic(err)
	}
}
