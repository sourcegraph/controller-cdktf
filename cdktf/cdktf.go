package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	hcversion "github.com/hashicorp/go-version"
	hcproduct "github.com/hashicorp/hc-install/product"
	tfreleases "github.com/hashicorp/hc-install/releases"
	cp "github.com/otiai10/copy"
	"github.com/sourcegraph/log"
	"github.com/sourcegraph/run"
	"github.com/sourcegraph/sourcegraph/lib/errors"
	"github.com/urfave/cli/v2"
	"sigs.k8s.io/yaml"

	"github.com/sourcegraph/controller-cdktf/internal/observability"
	"github.com/sourcegraph/controller-cdktf/internal/output"
)

func main() {
	liblog := observability.InitLogs("cdktf", "dev")
	defer liblog.Sync()

	sort.Sort(cli.CommandsByName(gen.Commands))
	sort.Sort(cli.FlagsByName(gen.Flags))

	if err := gen.Run(os.Args); err != nil {
		_ = output.Render(output.FormatText, err)
		os.Exit(1)
	}
}

type Config struct {
	Name     string      `json:"name"`
	Provider *SourceSpec `json:"provider"`
	Module   *SourceSpec `json:"module"`
}

type SourceSpec struct {
	// Name is the name of the provider or module
	// This field is only used internally to render the `cdktf.json` template
	// It will be ignored in the incoming config file
	Name string `json:"name,omitempty"`

	Source  string `json:"source"`
	Version string `json:"version,omitempty"`
}

type cdktfManifest struct {
	Language           string       `json:"language"`
	App                string       `json:"app"`
	SendCrashReports   bool         `json:"sendCrashReports"`
	TerraformProviders []SourceSpec `json:"terraformProviders,omitempty"`
	TerraformModules   []SourceSpec `json:"terraformModules,omitempty"`
	ProjectID          string       `json:"projectId"`
	Comment            string       `json:"//"`
}

var (
	configFlag = &cli.StringFlag{
		Name:     "config",
		Required: true,
	}
	moduleNameFlag = &cli.StringFlag{
		Name:     "moduleName",
		Usage:    "The go module name",
		Required: true,
	}
	packageNameFlag = &cli.StringFlag{
		Name:     "packageName",
		Usage:    "The go package name",
		Required: true,
	}
	outputFlag = &cli.StringFlag{
		Name:     "output",
		Usage:    "The output directory",
		Required: true,
	}

	//go:embed package.json
	packageJSONTemplateString string
	packageJSONTemplate       = template.Must(template.New("").Parse(packageJSONTemplateString))
)

type projectTemplateData struct {
	Config      Config
	PackageName string
	ModuleName  string
}

var gen = &cli.App{
	Name: "gen",
	Action: func(c *cli.Context) error {
		logger := log.Scoped("gen", "generate cdktf provider code")

		logger.Debug("current environment variables", log.Strings("env", os.Environ()))

		// workarounad for lack of well supported terraform toolchains for bazel
		// so we need to bring our own terraform and configure it in the path
		// so the cdktf-cli npm package can access it
		tfInstallDir, err := os.MkdirTemp("", "tf-bin")
		if err != nil {
			return errors.Wrap(err, "create temp tf-bin dir")
		}
		defer func() {
			_ = os.RemoveAll(tfInstallDir)
		}()
		installer := &tfreleases.ExactVersion{
			Product: hcproduct.Terraform,
			Version: hcversion.Must(hcversion.NewVersion("1.4.4")),
		}
		installer.InstallDir = tfInstallDir
		_, err = installer.Install(c.Context)
		if err != nil {
			return errors.Wrap(err, "install terraform")
		}
		_ = os.Setenv("PATH", tfInstallDir+string(os.PathListSeparator)+os.Getenv("PATH"))

		b, err := os.ReadFile(configFlag.Get(c))
		if err != nil {
			return errors.Wrap(err, "read config file")
		}

		var config Config
		if err := yaml.Unmarshal(b, &config); err != nil {
			return errors.Wrap(err, "unmarshal config file")
		}
		logger = logger.With(log.String("name", config.Name))
		if config.Provider != nil {
			logger = logger.With(log.String("provider", config.Provider.Source), log.String("version", config.Provider.Version))
		}
		if config.Module != nil {
			logger = logger.With(log.String("module", config.Module.Source), log.String("version", config.Module.Version))
		}

		if config.Provider != nil && config.Module != nil {
			return errors.New("provider and module can't be set at the same time")
		}

		m := cdktfManifest{
			Language:         "typescript",
			App:              "echo noop",
			SendCrashReports: false,
			ProjectID:        "noop",
		}
		if config.Provider != nil {
			// this is a special handling for provider name with hyphens
			providerName, ok := Last(strings.Split(config.Provider.Source, "/"))
			if !ok {
				return errors.Newf("provider name not found: %q", config.Provider.Source)
			}
			config.Provider.Name = providerName
			m.TerraformProviders = []SourceSpec{
				*config.Provider,
			}
		}
		if config.Module != nil {
			config.Module.Name = config.Name
			m.TerraformModules = []SourceSpec{
				*config.Module,
			}
		}
		var cdktfJSON bytes.Buffer
		enc := json.NewEncoder(&cdktfJSON)
		enc.SetEscapeHTML(false)
		if err := enc.Encode(m); err != nil {
			return errors.Wrap(err, "marshal cdktf.json")
		}
		logger.Debug("cdktf.json", log.String("cdktf.json", cdktfJSON.String()))

		data := projectTemplateData{
			Config:      config,
			PackageName: packageNameFlag.Get(c),
			ModuleName:  moduleNameFlag.Get(c),
		}
		var packageJSON bytes.Buffer
		if err := packageJSONTemplate.Execute(&packageJSON, data); err != nil {
			return errors.Wrap(err, "render package.json")
		}

		tmpDir, err := os.MkdirTemp("", "cdktf")
		if err != nil {
			return errors.Wrap(err, "create temp dir")
		}
		defer func() {
			_ = os.RemoveAll(tmpDir)
		}()
		logger = logger.With(log.String("tmpDir", tmpDir))

		logger.Debug("write package.json")
		if err := os.WriteFile(filepath.Join(tmpDir, "package.json"), packageJSON.Bytes(), 0644); err != nil {
			return errors.Wrap(err, "write package.json")
		}
		logger.Debug("write cdktf.json")
		if err := os.WriteFile(filepath.Join(tmpDir, "cdktf.json"), cdktfJSON.Bytes(), 0644); err != nil {
			return errors.Wrap(err, "write cdktf.json")
		}

		logger.Debug("compiling cdktf provider code")
		cmdCtx := observability.LogCommands(c.Context, logger)
		for _, cmd := range []string{
			"npm install --no-save",
			"npm run fetch",
			"npm run compile",
			"npm run pkg:go",
		} {
			if err := run.Cmd(cmdCtx, cmd).Dir(tmpDir).Run().Wait(); err != nil {
				return errors.Wrapf(err, "run: %q", cmd)
			}
		}

		cwd, err := os.Getwd()
		if err != nil {
			return errors.Wrap(err, "get working dir")
		}
		outputDir := filepath.Join(cwd, outputFlag.Get(c))
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return errors.Wrap(err, "create output dir")
		}
		if err := cp.Copy(filepath.Join(tmpDir, "dist", "go", config.Name), outputDir); err != nil {
			return errors.Wrap(err, "copy cdktf.out")
		}

		return nil
	},
	Flags: []cli.Flag{
		configFlag,
		moduleNameFlag,
		packageNameFlag,
		outputFlag,
	},
}

func Last[E any](s []E) (E, bool) {
	if len(s) == 0 {
		var zero E
		return zero, false
	}
	return s[len(s)-1], true
}
