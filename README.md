# cdktf-provider-generator

> EXPERIMENTAL

Generate your own pre-built CDKTF providers for Go.

CDKTF publishes pre-built providers but they do not follow the upstream providers version. Usually, you want to pin to a specific version to fit your own needs. This project aims to solve this problem by generating your own pre-built providers.

To do so, we replicate the setup from the CDKTF pre-built provider repo, e.g., https://github.com/cdktf/cdktf-provider-tfe, but use Bazel to generate all assets.

## Usage

If you're just looking to consume the generated code, all available providers are under `gen/` under its own Go module.

Use `hashicorp/tfe` provider:

```sh
go get github.com/controller-cdktf/gen/tfe
```

Use `hashicorp/google` provider:

```sh
go get github.com/controller-cdktf/gen/google
```

## Tasks

### Add a new provider

> NOTE: The name of the provider should always be one word and no special characters, e.g., use `googlebeta` instead of `google-beta`.

For example, let's add the `kubernetes` provider:

Create a config file at `config/kubernetes.yaml`

```yaml
name: kubernetes
provider:
  source: registry.terraform.io/hashicorp/kubernetes
  version: ~> 2.15.0
```

Add a new target in `BUILD.bazel`

```diff
--- a/BUILD.bazel
+++ b/BUILD.bazel

+
+cdktf(
+    name = "kubernetes",
+    config = "kubernetes.yaml",
+    module_name = module_name,
+)
```

Generate the code

```sh
bazel run //:kubernetes
bazel run //:gazelle
```

Finally, commit all generated code and open a pull request.

### Add a new module

For example, let's add the `awsvpc` module:

Create a config file at `config/awsvpc.yaml`

```yaml
name: awsvpc
module:
  source: terraform-aws-modules/vpc/aws
  version: 3.19.0
```

Add a new target in `BUILD.bazel`

```diff
--- a/BUILD.bazel
+++ b/BUILD.bazel

+
+cdktf(
+    name = "awsvpc",
+    config = "awsvpc.yaml",
+    module_name = module_name,
+)
```

Generate the code

```sh
bazel run //:awsvpc
bazel run //:gazelle
```

Finally, commit all generated code and open a pull request.

### Update a provider or module version

First, update the version in the corresponding config file, e.g., `gen/kubernetes.yaml`.

```diff
--- a/config/kubernetes.yaml
+++ b/config/kubernetes.yaml
@@ -1,4 +1,4 @@
 name: kubernetes
 provider:
   source: registry.terraform.io/hashicorp/kubernetes
-  version: 2.15.0
+  version: 2.29.0
```

Then, generate the code:

```sh
bazel run //:kubernetes
bazel run //:gazelle
```

Finally, commit all generated code and open a pull request.

## Roadmap

- [ ] Add support to configure target `cdktf` and `jsii` version
- [ ] Add support to configure target Terraform version

## FAQ

### Why not use `cdktf get`?

With a centralized `cdktf.json` where you declare all the providers and modules, changes to the file will require re-generating all providers and modules, which is slow.

By utilizing Bazel, we can only generate the code for the provider or module that has changed with its caching strategy.

### Weird things happen during generation with Bazel?

First, try to clean the Bazel cache:

```sh
bazel clean --expunge
```

Then, try again.
