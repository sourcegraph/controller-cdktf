package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3Bucket struct {
	// aws_account block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/data_loss_prevention_discovery_config#aws_account DataLossPreventionDiscoveryConfig#aws_account}
	AwsAccount *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketAwsAccount `field:"optional" json:"awsAccount" yaml:"awsAccount"`
	// The bucket name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/data_loss_prevention_discovery_config#bucket_name DataLossPreventionDiscoveryConfig#bucket_name}
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
}

