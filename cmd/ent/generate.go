package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/spf13/cobra"
	"github.com/walkersumida/go-api-server/internal/pkg/path"
)

var cmdGenerate = &cobra.Command{
	Use:   "generate",
	Short: "Generate ent codes",
	Run: func(cmd *cobra.Command, args []string) {
		generate()
	},
}

func generate() {
	p := path.Project()
	err := entc.Generate(p+"/ent/schema", &gen.Config{
		Features: []gen.Feature{gen.FeatureVersionedMigration},
	})
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
