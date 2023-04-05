"Bazel go dependencies"

load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_dependencies():
    """The go dependencies in this macro are auto-updated by gazelle
    To update run,
        bazel run //:gazelle-update-repos
    """
    pass
    go_repository(
        name = "cc_mvdan_gofumpt",
        build_file_proto_mode = "disable_global",
        importpath = "mvdan.cc/gofumpt",
        sum = "h1:JVf4NN1mIpHogBj7ABpgOyZc65/UUOkKQFkoURsz4MM=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_antonmedv_expr",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/antonmedv/expr",
        sum = "h1:j4HI3NHEdgDnN9p6oI6Ndr0G5QryMY0FNxT4ONrFDGU=",
        version = "v1.9.0",
    )
    go_repository(
        name = "com_github_aws_constructs_go_constructs_v10",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/aws/constructs-go/constructs/v10",
        sum = "h1:Ys+0FJmsRPnR2oqcozlqcl2IHPSo5BQoCaOk0mq6LK8=",
        version = "v10.1.244",
    )
    go_repository(
        name = "com_github_aws_jsii_runtime_go",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/aws/jsii-runtime-go",
        sum = "h1:8Ol3FA5PaBbJbCKTQhrqfImUHsfm3ClxFc5011fqjZo=",
        version = "v1.74.0",
    )
    go_repository(
        name = "com_github_davecgh_go_spew",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/davecgh/go-spew",
        sum = "h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_djherbis_buffer",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/djherbis/buffer",
        sum = "h1:PH5Dd2ss0C7CRRhQCZ2u7MssF+No9ide8Ye71nPHcrQ=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_djherbis_nio_v3",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/djherbis/nio/v3",
        sum = "h1:6wxhnuppteMa6RHA4L81Dq7ThkZH8SwnDzXDYy95vB4=",
        version = "v3.0.1",
    )
    go_repository(
        name = "com_github_fatih_color",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/fatih/color",
        sum = "h1:8LOYc1KYPPmyKMuN8QV2DNRWNbLo6LZ0iLs8+mlH53w=",
        version = "v1.13.0",
    )
    go_repository(
        name = "com_github_frankban_quicktest",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/frankban/quicktest",
        sum = "h1:FJKSZTDHjyhriyC81FLQ0LY93eSai0ZyR/ZIkd3ZUKE=",
        version = "v1.14.3",
    )
    go_repository(
        name = "com_github_go_logr_logr",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/go-logr/logr",
        sum = "h1:2DntVwHkVopvECVRSlL5PSo9eG+cAkDCuckLubN+rq0=",
        version = "v1.2.3",
    )
    go_repository(
        name = "com_github_go_logr_stdr",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/go-logr/stdr",
        sum = "h1:hSWxHoqTgW2S2qGc0LTAI563KZ5YKYRhT3MFKZMbjag=",
        version = "v1.2.2",
    )
    go_repository(
        name = "com_github_google_go_cmp",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/google/go-cmp",
        sum = "h1:O2Tfq5qg4qc4AmwVlvv0oLiVAGB7enBSJ2x2DqQFi38=",
        version = "v0.5.9",
    )
    go_repository(
        name = "com_github_hashicorp_terraform_cdk_go_cdktf",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/hashicorp/terraform-cdk-go/cdktf",
        sum = "h1:QY2238bD3bGIZRILpvppMKAWQd77zINKktVVzH3dSUk=",
        version = "v0.13.3",
    )
    go_repository(
        name = "com_github_hexops_autogold",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/hexops/autogold",
        sum = "h1:YgxF9OHWbEIUjhDbpnLhgVsjUDsiHDTyDfy2lrfdlzo=",
        version = "v1.3.1",
    )
    go_repository(
        name = "com_github_hexops_gotextdiff",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/hexops/gotextdiff",
        sum = "h1:gitA9+qJrrTCsiCl7+kh75nPqQt1cx4ZkudSTLoUqJM=",
        version = "v1.0.3",
    )
    go_repository(
        name = "com_github_hexops_valast",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/hexops/valast",
        sum = "h1:oBoGERMJh6UZdRc6cduE1CTPK+VAdXA59Y1HFgu3sm0=",
        version = "v1.4.3",
    )
    go_repository(
        name = "com_github_itchyny_gojq",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/itchyny/gojq",
        sum = "h1:YhLueoHhHiN4mkfM+3AyJV6EPcCxKZsOnYf+aVSwaQw=",
        version = "v0.12.11",
    )
    go_repository(
        name = "com_github_itchyny_timefmt_go",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/itchyny/timefmt-go",
        sum = "h1:G0INE2la8S6ru/ZI5JecgyzbbJNs5lG1RcBqa7Jm6GE=",
        version = "v0.1.5",
    )
    go_repository(
        name = "com_github_kr_pretty",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/kr/pretty",
        sum = "h1:WgNl7dwNpEZ6jJ9k1snq4pZsg7DOEN8hP9Xw0Tsjwk0=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_kr_text",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/kr/text",
        sum = "h1:5Nx0Ya0ZqY2ygV366QzturHI13Jq95ApcVaJBhpS+AY=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_masterminds_semver_v3",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/Masterminds/semver/v3",
        sum = "h1:3MEsd0SM6jqZojhjLWWeBY+Kcjy9i6MQAeY7YgDP83g=",
        version = "v3.2.0",
    )
    go_repository(
        name = "com_github_mattn_go_colorable",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/mattn/go-colorable",
        sum = "h1:fFA4WZxdEF4tXPZVKMLwD8oUnCTTo08duU7wxecdEvA=",
        version = "v0.1.13",
    )
    go_repository(
        name = "com_github_mattn_go_isatty",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/mattn/go-isatty",
        sum = "h1:bq3VjFmv/sOjHtdEhmkEV4x1AJtvUvOJ2PFAZ5+peKQ=",
        version = "v0.0.16",
    )
    go_repository(
        name = "com_github_mattn_go_runewidth",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/mattn/go-runewidth",
        sum = "h1:+xnbZSEeDbOIg5/mE6JF0w6n9duR1l3/WmbinWVwUuU=",
        version = "v0.0.14",
    )
    go_repository(
        name = "com_github_nightlyone_lockfile",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/nightlyone/lockfile",
        sum = "h1:RHep2cFKK4PonZJDdEl4GmkabuhbsRMgk/k3uAmxBiA=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_olekukonko_tablewriter",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/olekukonko/tablewriter",
        sum = "h1:P2Ga83D34wi1o9J6Wh1mRuqd4mF/x/lgBS7N7AbDhec=",
        version = "v0.0.5",
    )
    go_repository(
        name = "com_github_pmezard_go_difflib",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/pmezard/go-difflib",
        sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_rivo_uniseg",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/rivo/uniseg",
        sum = "h1:S1pD9weZBuJdFmowNwbpi7BJ8TNftyUImj/0WQi72jY=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_rogpeppe_go_internal",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/rogpeppe/go-internal",
        sum = "h1:/FiVV8dS/e+YqF2JvO3yXRFbBLTIuSDkuC7aBOAvL+k=",
        version = "v1.6.1",
    )
    go_repository(
        name = "com_github_sourcegraph_run",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/sourcegraph/run",
        sum = "h1:3A8w5e8HIYPfafHekvmdmmh42RHKGVhmiTZAPJclg7I=",
        version = "v0.12.0",
    )
    go_repository(
        name = "com_github_spf13_pflag",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/spf13/pflag",
        sum = "h1:iy+VFUOCP1a+8yFto/drg2CJ5u0yRoB7fZw3DKv/JXA=",
        version = "v1.0.5",
    )
    go_repository(
        name = "com_github_stretchr_objx",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/stretchr/objx",
        sum = "h1:M2gUjqZET1qApGOWNSnZ49BAIMX4F/1plDv3+l31EJ4=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_stretchr_testify",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/stretchr/testify",
        sum = "h1:w7B6lhMri9wdJUVmEZPGGhZzrYTPvgJArz7wNPgYKsk=",
        version = "v1.8.1",
    )
    go_repository(
        name = "com_github_yuin_goldmark",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/yuin/goldmark",
        sum = "h1:fVcFKWvrslecOb/tg+Cc05dkeYx540o0FuFt3nUVDoE=",
        version = "v1.4.13",
    )
    go_repository(
        name = "dev_bobheadxi_go_gobenchdata",
        build_file_proto_mode = "disable_global",
        importpath = "go.bobheadxi.dev/gobenchdata",
        sum = "h1:6YMtFDViJ9mLs1WnpYN74+VgtOhsT9cZbrd7OCBpn30=",
        version = "v1.3.0",
    )
    go_repository(
        name = "dev_bobheadxi_go_streamline",
        build_file_proto_mode = "disable_global",
        importpath = "go.bobheadxi.dev/streamline",
        sum = "h1:IqKSA1TbeuDqCzYNAwtlh8sqf3tsQus8XgJdkCWFT8c=",
        version = "v1.2.1",
    )
    go_repository(
        name = "in_gopkg_check_v1",
        build_file_proto_mode = "disable_global",
        importpath = "gopkg.in/check.v1",
        sum = "h1:yhCVgyC4o1eVCa2tZl7eS0r+SDo693bJlVdllGtEeKM=",
        version = "v0.0.0-20161208181325-20d25e280405",
    )
    go_repository(
        name = "in_gopkg_yaml_v2",
        build_file_proto_mode = "disable_global",
        importpath = "gopkg.in/yaml.v2",
        sum = "h1:D8xgwECY7CYvx+Y2n4sBz93Jn9JRvxdiyyo8CTfuKaY=",
        version = "v2.4.0",
    )
    go_repository(
        name = "in_gopkg_yaml_v3",
        build_file_proto_mode = "disable_global",
        importpath = "gopkg.in/yaml.v3",
        sum = "h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=",
        version = "v3.0.1",
    )
    go_repository(
        name = "io_opentelemetry_go_otel",
        build_file_proto_mode = "disable_global",
        importpath = "go.opentelemetry.io/otel",
        sum = "h1:kfToEGMDq6TrVrJ9Vht84Y8y9enykSZzDDZglV0kIEk=",
        version = "v1.11.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_sdk",
        build_file_proto_mode = "disable_global",
        importpath = "go.opentelemetry.io/otel/sdk",
        sum = "h1:ZnKIL9V9Ztaq+ME43IUi/eo22mNsb6a7tGfzaOWB5fo=",
        version = "v1.11.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_trace",
        build_file_proto_mode = "disable_global",
        importpath = "go.opentelemetry.io/otel/trace",
        sum = "h1:20U/Vj42SX+mASlXLmSGBg6jpI1jQtv682lZtTAOVFI=",
        version = "v1.11.0",
    )
    go_repository(
        name = "org_bitbucket_creachadair_shell",
        build_file_proto_mode = "disable_global",
        importpath = "bitbucket.org/creachadair/shell",
        sum = "h1:Z96pB6DkSb7F3Y3BBnJeOZH2gazyMTWlvecSD4vDqfk=",
        version = "v0.0.7",
    )
    go_repository(
        name = "org_golang_x_crypto",
        build_file_proto_mode = "disable_global",
        importpath = "golang.org/x/crypto",
        sum = "h1:7I4JAnoQBe7ZtJcBaYHi5UtiO8tQHbUSXxL+pnGRANg=",
        version = "v0.0.0-20210921155107-089bfa567519",
    )
    go_repository(
        name = "org_golang_x_lint",
        build_file_proto_mode = "disable_global",
        importpath = "golang.org/x/lint",
        sum = "h1:VLliZ0d+/avPrXXH+OakdXhpJuEoBZuwh1m2j7U6Iug=",
        version = "v0.0.0-20210508222113-6edffad5e616",
    )
    go_repository(
        name = "org_golang_x_mod",
        build_file_proto_mode = "disable_global",
        importpath = "golang.org/x/mod",
        sum = "h1:LapD9S96VoQRhi/GrNTqeBJFrUjs5UHCAtTlgwA5oZA=",
        version = "v0.7.0",
    )
    go_repository(
        name = "org_golang_x_net",
        build_file_proto_mode = "disable_global",
        importpath = "golang.org/x/net",
        sum = "h1:GyT4nK/YDHSqa1c4753ouYCDajOYKTja9Xb/OHtgvSw=",
        version = "v0.5.0",
    )
    go_repository(
        name = "org_golang_x_sync",
        build_file_proto_mode = "disable_global",
        importpath = "golang.org/x/sync",
        sum = "h1:wsuoTGHzEhffawBOhz5CYhcrV4IdKZbEyZjBMuTp12o=",
        version = "v0.1.0",
    )
    go_repository(
        name = "org_golang_x_sys",
        build_file_proto_mode = "disable_global",
        importpath = "golang.org/x/sys",
        sum = "h1:Zr2JFtRQNX3BCZ8YtxRE9hNJYC8J6I1MVbMg6owUp18=",
        version = "v0.4.0",
    )
    go_repository(
        name = "org_golang_x_term",
        build_file_proto_mode = "disable_global",
        importpath = "golang.org/x/term",
        sum = "h1:JGgROgKl9N8DuW20oFS5gxc+lE67/N3FcwmBPMe7ArY=",
        version = "v0.0.0-20210927222741-03fcf44c2211",
    )
    go_repository(
        name = "org_golang_x_text",
        build_file_proto_mode = "disable_global",
        importpath = "golang.org/x/text",
        sum = "h1:olpwvP2KacW1ZWvsR7uQhoyTYvKAupfQrRGBFM352Gk=",
        version = "v0.3.7",
    )
    go_repository(
        name = "org_golang_x_tools",
        build_file_proto_mode = "disable_global",
        importpath = "golang.org/x/tools",
        sum = "h1:+bSpV5HIeWkuvgaMfI3UmKRThoTA5ODJTUd8T17NO+4=",
        version = "v0.5.0",
    )
    go_repository(
        name = "org_golang_x_xerrors",
        build_file_proto_mode = "disable_global",
        importpath = "golang.org/x/xerrors",
        sum = "h1:E7g+9GITq07hpfrRu66IVDexMakfv52eLZ2CXBWiKr4=",
        version = "v0.0.0-20191204190536-9bdfabe68543",
    )
