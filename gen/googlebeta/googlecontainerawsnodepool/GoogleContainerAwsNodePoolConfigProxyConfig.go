package googlecontainerawsnodepool


type GoogleContainerAwsNodePoolConfigProxyConfig struct {
	// The ARN of the AWS Secret Manager secret that contains the HTTP(S) proxy configuration.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_aws_node_pool#secret_arn GoogleContainerAwsNodePool#secret_arn}
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
	// The version string of the AWS Secret Manager secret that contains the HTTP(S) proxy configuration.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_aws_node_pool#secret_version GoogleContainerAwsNodePool#secret_version}
	SecretVersion *string `field:"required" json:"secretVersion" yaml:"secretVersion"`
}

