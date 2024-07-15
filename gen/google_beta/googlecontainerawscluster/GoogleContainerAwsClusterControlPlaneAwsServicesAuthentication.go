package googlecontainerawscluster


type GoogleContainerAwsClusterControlPlaneAwsServicesAuthentication struct {
	// The Amazon Resource Name (ARN) of the role that the Anthos Multi-Cloud API will assume when managing AWS resources on your account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_aws_cluster#role_arn GoogleContainerAwsCluster#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Optional. An identifier for the assumed role session. When unspecified, it defaults to `multicloud-service-agent`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_aws_cluster#role_session_name GoogleContainerAwsCluster#role_session_name}
	RoleSessionName *string `field:"optional" json:"roleSessionName" yaml:"roleSessionName"`
}

