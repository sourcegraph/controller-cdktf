# AGENTS.md

Guidance for coding agents working in `controller-cdktf`.

## What this repo is

This repo holds **generated** [cdktf](https://github.com/hashicorp/terraform-cdk)
Go code, consumed internally by [sourcegraph/controller]. The generated packages
live under `gen/<provider>/` (e.g. `gen/google`, `gen/aws`, `gen/kubernetes`) and
are checked into Git on purpose.

Do not hand-edit anything under `gen/` — it is regenerated. The implementation
that uses these packages lives in [sourcegraph/controller], and the generator is
[sourcegraph/cdktf-provider-gen].

## Layout

- `*.yml` (e.g. `google.yml`, `aws.yml`) — one config per provider/module. Each
  declares the provider `source`, `version`, target Go `packageName`, and
  `output: gen`.
- `gen/<provider>/` — generated Go code, plus `gen/versions.json` and
  `gen/constraints.json`.
- `tools/tools.go` — pins the `cdktf-provider-gen` generator binary as a tool
  dependency.
- `Makefile` — drives generation.

## Toolchain

Versions are pinned in `.tool-versions` and `go.mod`:

- Go (module declares `go 1.25.0`; `.tool-versions` lists `golang 1.20`)
- Terraform `1.2.6`, Node.js `20.13.1`
- CDKTF version is set by `CDKTF_VERSION` in the `Makefile`

## Common commands

```bash
# Build the generator into .bin/
make install

# List all generatable targets (derived from *.yml files)
make targets

# Generate a single provider/module (target name == <name>.yml without extension)
make google

# Generate everything (install + all providers), in parallel
make -j4
```

## Conventions

- To add a provider/module, add a `<name>.yml` config (see existing files) and
  run `make <name>`. Follow [sourcegraph/cdktf-provider-gen#usage].
- To upgrade CDKTF, review the upstream changelog, bump `CDKTF_VERSION` in the
  `Makefile`, and re-run `make -j4`.
- Providers are split into multiple modules because the combined generated code
  exceeds `go get` limits.
- `gen/*/**` is marked in `.gitattributes`; `.gitignore` keeps `.bin/` out of Git
  but retains `gen/*/jsii/*.tgz`.

[sourcegraph/controller]: https://github.com/sourcegraph/controller
[sourcegraph/cdktf-provider-gen]: https://github.com/sourcegraph/cdktf-provider-gen
[sourcegraph/cdktf-provider-gen#usage]: https://github.com/sourcegraph/cdktf-provider-gen#usage
