"""
Rules to generate cdktf providers code
"""

load("@bazel_skylib//lib:paths.bzl", "paths")
load("@bazel_skylib//rules:write_file.bzl", "write_file")

def cdktf(name, config, module_prefix, out_pkg):
    """
    Generates a cdktf provider code from a config file

    - name: aws
    - module_prefix: github.com/hashicorp/terraform-cdk-providers
    - out_pkg: gen

    the result go module name will be github.com/hashicorp/terraform-cdk-providers/gen/aws

    Args:
        name: name of the rule, this will be the package name
        config: path to the config file
        module_prefix: prefix of the go module to generate
        out_pkg: name of the package to generate under the parent module
    """

    native.genrule(
        name = "{name}__get".format(name = name),
        srcs = [config],
        outs = ["{name}__gen".format(name = name)],
        cmd = " && ".join([
                "export GO_REL_PATH=`dirname $(location @go_sdk//:bin/go)`",
                "export GO_ABS_PATH=`cd $$GO_REL_PATH && pwd`",
                "export NPM_REL_PATH=`dirname $(location @nodejs//:npm)`",
                "export NPM_ABS_PATH=`cd $$NPM_REL_PATH && pwd`",
                "export PATH=$$GO_ABS_PATH:$$NPM_ABS_PATH:$$PATH",
                "export HOME=$(GENDIR)",
                "export GOPATH=/nonexist-gopath",
                "export TF_PLUGIN_CACHE_DIR=$$(mktemp -d || mktemp -d -t terraform-plugin-cache-dir)",
                "export npm_config_cache=$$(mktemp -d || mktemp -d -t npm-cache-dir)",
                "export GOMODCACHE=$$(mktemp -d || mktemp -d -t gomodcache-dir)",
                "export SRC_LOG_LEVEL=dbug",
                "$(location //cdktf:cdktf) --output $@ --config $< --packageName {name} --moduleName {module_prefix}/{out_pkg}".format(name = name, module_prefix = module_prefix, out_pkg = out_pkg),
            ]),
        toolchains = [
            "@nodejs_toolchains//:resolved_toolchain",
        ],
        tools = [
            "@nodejs_toolchains//:resolved_toolchain",
            "@go_sdk//:bin/go",
            "@nodejs//:npm",
            "//cdktf:cdktf",
        ],
    )

    # copy generate code into the workspace
    # https://www.aspect.dev/blog/bazel-can-write-to-the-source-folder
    bazel_bin_out_dir = paths.join("bazel-bin", native.package_name(), "{name}__gen".format(name = name))
    workspace_out_dir = paths.join(native.package_name(), out_pkg, name)
    write_file(
        name = "{name}__gen_copy".format(name = name),
        out =  "{name}__gen_copy.sh".format(name = name),
        content = [
            # This depends on bash, would need tweaks for Windows
            "#!/usr/bin/env bash",
            # Bazel gives us a way to access the source folder!
            "cd $BUILD_WORKSPACE_DIRECTORY",
        ] + [
            # Paths are now relative to the workspace.
            # We can copy files from bazel-bin to the sources
            "rm -rf {0}".format(workspace_out_dir),
            "mkdir -p {0}".format(workspace_out_dir),
            "cp -r {0}/. {1}/".format(
                bazel_bin_out_dir,
                workspace_out_dir,
            )
        ],
    )

    # `bazel run` will run this script to copy the generated providers code into the workspace
    native.sh_binary(
        name = name,
        srcs = ["{name}__gen_copy.sh".format(name = name)],
        # data = [":{name}__gen".format(name = name)],
        data = [":{name}__gen/.".format(name = name)],
    )
