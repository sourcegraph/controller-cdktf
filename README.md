# controller-cdktf

This repo contains the generated code for [cdktf](https://github.com/hashicorp/terraform-cdk) in Go. The implementation is being developed in [sourcegraph/controller].

This package is only used internally at Sourcegraph - the generated code is public for ease of use and avoid performance issue with large amount of generated content being tracked in a Git repository.

## Development

### Adding a new provider or module to CDKTF

Edit [`cdktf.json`](./cdktf.json), and make sure modules and providers are pinned to a specific version. e.g.

```json
{
  "terraformProviders": [
    {
      "name": "google",
      "source": "hashicorp/google",
      "version": "~> 4.38.0"
    }
  ],
  "terraformModules": [
    {
      "name": "googlesqldb",
      "source": "git::https://github.com/michaellzc/terraform-google-sql-db//modules/postgresql?ref=feat/support-dynamic-iam-users"
    }
  ]
}
```

Run the command below to auto-gen codes for added providers or modules:

```sh
go generate .
```

Commit all generated changes.

### Upgrading CDKTF

Review the [changelog](https://developer.hashicorp.com/terraform/cdktf/release#upgrade-guides) of the target release.
Watch out for breaking changes and adjust the upgrade plan if neccessary.

Bump the version in [`gen.go`](./gen.go):

> no `v` prefix

```diff
--- a/gen.go
+++ b/gen.go
@@ -1,3 +1,3 @@
 package cdktf

-//go:generate go run . --version 0.13.3
+//go:generate go run . --version 0.13.0
```

Re-generate the providers and modules code:

```sh
go generate ./
```

Commit all changes in this repo. Upon merging the PR, then open another pull request in [sourcegraph/controller] to upgrade all `github.com/sourcegraph/controller-cdktf/gen/*` packages:

```sh
go get -u github.com/sourcegraph/controller-cdktf/gen/...
```

## FAQ

### Why not use the pre-built providers?

We would like to use specific versions of provides, hence we are not using pre-built providers, such as [cdktf/cdktf-provider-google-go](https://github.com/cdktf/cdktf-provider-google-go).

### Why multiple modules instead of one?

Generated code combined from all providers and modules exceed the limit of `go get`, and this cannot be changed.

[sourcegraph/controller]: https://github.com/sourcegraph/controller
