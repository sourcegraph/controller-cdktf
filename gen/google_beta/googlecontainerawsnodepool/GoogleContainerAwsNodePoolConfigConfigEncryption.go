package googlecontainerawsnodepool


type GoogleContainerAwsNodePoolConfigConfigEncryption struct {
	// The ARN of the AWS KMS key used to encrypt node pool configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_aws_node_pool#kms_key_arn GoogleContainerAwsNodePool#kms_key_arn}
	KmsKeyArn *string `field:"required" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

