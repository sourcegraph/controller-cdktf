package googledatalosspreventiondiscoveryconfig


type GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3Bucket struct {
	// aws_account block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_data_loss_prevention_discovery_config#aws_account GoogleDataLossPreventionDiscoveryConfig#aws_account}
	AwsAccount *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceAmazonS3BucketAwsAccount `field:"optional" json:"awsAccount" yaml:"awsAccount"`
	// The bucket name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_data_loss_prevention_discovery_config#bucket_name GoogleDataLossPreventionDiscoveryConfig#bucket_name}
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
}

