package googlecontainerawscluster


type GoogleContainerAwsClusterControlPlaneProxyConfig struct {
	// The ARN of the AWS Secret Manager secret that contains the HTTP(S) proxy configuration.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_aws_cluster#secret_arn GoogleContainerAwsCluster#secret_arn}
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
	// The version string of the AWS Secret Manager secret that contains the HTTP(S) proxy configuration.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_aws_cluster#secret_version GoogleContainerAwsCluster#secret_version}
	SecretVersion *string `field:"required" json:"secretVersion" yaml:"secretVersion"`
}

