package googlecontainerawsnodepool


type GoogleContainerAwsNodePoolConfigRootVolume struct {
	// Optional. The number of I/O operations per second (IOPS) to provision for GP3 volume.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_aws_node_pool#iops GoogleContainerAwsNodePool#iops}
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Optional.
	//
	// The Amazon Resource Name (ARN) of the Customer Managed Key (CMK) used to encrypt AWS EBS volumes. If not specified, the default Amazon managed key associated to the AWS region where this cluster runs will be used.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_aws_node_pool#kms_key_arn GoogleContainerAwsNodePool#kms_key_arn}
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Optional.
	//
	// The size of the volume, in GiBs. When unspecified, a default value is provided. See the specific reference in the parent resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_aws_node_pool#size_gib GoogleContainerAwsNodePool#size_gib}
	SizeGib *float64 `field:"optional" json:"sizeGib" yaml:"sizeGib"`
	// Optional. Type of the EBS volume. When unspecified, it defaults to GP2 volume. Possible values: VOLUME_TYPE_UNSPECIFIED, GP2, GP3.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_aws_node_pool#volume_type GoogleContainerAwsNodePool#volume_type}
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}
