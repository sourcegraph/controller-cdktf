# controller-cdktf

This repo contains the generated code for [cdktf](https://github.com/hashicorp/terraform-cdk) in Go. The implementation is being developed in sourcegraph/controller.

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

Upgrade `cdktf` Go client:

```sh
go get github.com/hashicorp/terraform-cdk-go/cdktf@v$VERSION
go mod tidy
```

Bump the version in [`gen.go`](./gen.go) to upgrade `cdktf-cli`:

```diff
--- a/gen.go
+++ b/gen.go
@@ -1,3 +1,3 @@
 package cdktf

-//go:generate npx --yes cdktf-cli@0.13.0 get -o gen
+//go:generate npx --yes cdktf-cli@0.13.3 get -o gen
```

Re-genenerate the providers and modules code:

```sh
go generate ./
```

Commit all changes.
