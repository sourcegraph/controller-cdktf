package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/sourcegraph/run"
)

const (
	outputDir = "gen"
)

var (
	version = flag.String("version", "", "target cdktf version, no 'v' prefix, e.g. '0.13.3'")
)

//go:generate go run . --version 0.13.3
func main() {
	flag.Parse()

	if *version == "" {
		log.Fatal("no --version specified")
	}
	if strings.HasPrefix(*version, "v") {
		log.Fatalf("version must not start with 'v', %q", *version)
	}

	ctx := context.Background()

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	if err := run.Cmd(ctx, "npx", "--yes", "cdktf-cli@"+*version, "get", "--output", outputDir).Run().Stream(os.Stderr); err != nil {
		log.Fatal(err)
	}

	files, err := os.ReadDir(outputDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			path := filepath.Join(cwd, outputDir, file.Name())
			if err := run.Cmd(ctx, "go", "mod", "init", pkgName(path)).Dir(path).Run().Stream(os.Stdout); err != nil {
				log.Fatal(err)
			}
			if err := run.Cmd(ctx, "go", "get", "github.com/hashicorp/terraform-cdk-go/cdktf@v"+*version).Dir(path).Run().Stream(os.Stdout); err != nil {
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
	return fmt.Sprintf("github.com/sourcegraph/controller-cdktf/%s/%s", outputDir, filepath.Base(path))
}
