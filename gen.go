package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/sourcegraph/run"
)

var version = flag.String("version", "v0.13.3", "target cdktf version")

//go:generate go run .
func main() {
	flag.Parse()
	ctx := context.Background()

	if err := run.Cmd(ctx, "npx", "--yes", "cdktf-cli@"+*version, "get", "-o", "gen"); err != nil {
		log.Fatal(err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	files, err := os.ReadDir("gen")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			path := filepath.Join(cwd, "gen", file.Name())
			if err := run.Cmd(ctx, "go", "mod", "init", pkgName(path)).Dir(path).Run().Stream(os.Stdout); err != nil {
				log.Fatal(err)
			}
			if err := run.Cmd(ctx, "go", "get", "github.com/hashicorp/terraform-cdk-go/cdktf@"+*version).Dir(path).Run().Stream(os.Stdout); err != nil {
				log.Fatal(err)
			}
			if err := run.Cmd(ctx, "go", "mod", "tidy").Dir(path).Run().Stream(os.Stdout); err != nil {
				log.Fatal(err)
			}
		}
	}
}

// pkgName returns the package name for the given path and remove
// any underscores.
func pkgName(path string) string {
	return fmt.Sprintf("github.com/sourcegraph/controller-cdktf/gen/%s", filepath.Base(path))
}
