package main

import (
	"log"
	"os"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/spf13/cobra"
	"github.com/walkersumida/go-api-server/internal/pkg/path"
)

var excludePaths = []string{"schema"}

var cmdGenerate = &cobra.Command{
	Use:   "generate",
	Short: "Generate ent codes",
	Run: func(cmd *cobra.Command, args []string) {
		generate()
	},
}

func generate() {
	p := path.Project()
	dirs, err := os.ReadDir(p + "/ent")
	if err != nil {
		log.Fatalf("failed to read dir: %v", err)
	}

	for _, dir := range dirs {
		if isExcludePath(dir.Name()) {
			continue
		}
		if dir.IsDir() {
			err := os.RemoveAll(p + "/ent/" + dir.Name())
			if err != nil {
				log.Fatalf("failed to remove dir: %v", err)
			}
			continue
		}
		os.Remove(p + "/ent/" + dir.Name())
	}

	err = entc.Generate(p+"/ent/schema", &gen.Config{
		Features: []gen.Feature{gen.FeatureVersionedMigration},
	})
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}

func isExcludePath(path string) bool {
	for _, p := range excludePaths {
		if p == path {
			return true
		}
	}
	return false
}
