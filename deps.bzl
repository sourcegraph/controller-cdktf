load("@bazel_gazelle//:deps.bzl", "go_repository")

"Bazel go dependencies"

def go_dependencies():
    """The go dependencies in this macro are auto-updated by gazelle
    To update run,
        bazel run //:gazelle-update-repos
    """
    pass
    go_repository(
        name = "com_github_aws_constructs_go_constructs_v10",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/aws/constructs-go/constructs/v10",
        sum = "h1:XW89VuKlwwzgU77Nyzj30SPn2SFtDYkzfF+QN84b5bQ=",
        version = "v10.2.69",
    )
    go_repository(
        name = "com_github_aws_jsii_runtime_go",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/aws/jsii-runtime-go",
        sum = "h1:C8BMeDgDz4IcDvgKFPMDt2OUGsTiANVu/YZGtwUoJqE=",
        version = "v1.87.0",
    )
    go_repository(
        name = "com_github_hashicorp_terraform_cdk_go_cdktf",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/hashicorp/terraform-cdk-go/cdktf",
        sum = "h1:rpgeX9zQ9jhHoBc/qKEoVOAP/iVhwPN463UyhZKTwy0=",
        version = "v0.16.3",
    )
