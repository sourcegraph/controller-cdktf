# controller-cdktf

This repo contains the generated Go code for [CDK-Terrain](https://github.com/open-constructs/cdk-terrain), the open source successor to CDK for Terraform (CDKTF). The implementation is being developed in [sourcegraph/controller]. The repository name and generated Go module paths retain `cdktf` for compatibility with existing consumers.

This package is only used internally at Sourcegraph - the generated code is public for ease of use and avoid performance issue with large amount of generated content being tracked in a Git repository.

## Usage

### Adding a new provider or module to CDK-Terrain

Follow https://github.com/sourcegraph/cdktf-provider-gen#usage

```bash
make <target>
```

### Upgrading CDK-Terrain

Review the [CDK-Terrain changelog](https://github.com/open-constructs/cdk-terrain/blob/main/CHANGELOG.md) of the target release.
Watch out for breaking changes and adjust the upgrade plan if neccessary.

Bump `CDKTN_VERSION` in `Makefile`:

```diff
-CDKTN_VERSION=0.24.0
+CDKTN_VERSION=0.25.0
```

Re-generate all providers and modules:

```bash
make -j4
```

If you upgrade google terraform provider, follow the version upgrade guide eg. [v6 upgrade](https://github.com/hashicorp/terraform-provider-google/blob/main/website/docs/guides/version_6_upgrade.html.markdown) and update the controller.

## FAQ

### Why not use the pre-built providers?

We would like to use specific versions of provides, hence we are not using pre-built providers, such as [cdktf/cdktf-provider-google-go](https://github.com/cdktf/cdktf-provider-google-go).

### Why multiple modules instead of one?

Generated code combined from all providers and modules exceed the limit of `go get`, and this cannot be changed.

[sourcegraph/controller]: https://github.com/sourcegraph/controller
