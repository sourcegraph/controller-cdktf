for module in ./gen/* ; do
    pushd $module
    go mod init github.com/sourcegraph/controller-cdktf/gen/$(basename $module)
    go get github.com/hashicorp/terraform-cdk-go/cdktf@v0.13.3
    go mod tidy
    popd
done
