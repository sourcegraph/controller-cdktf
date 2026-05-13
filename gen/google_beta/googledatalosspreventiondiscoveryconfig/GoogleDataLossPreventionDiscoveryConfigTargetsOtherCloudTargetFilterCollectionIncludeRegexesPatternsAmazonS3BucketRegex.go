package googledatalosspreventiondiscoveryconfig


type GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegex struct {
	// aws_account_regex block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_data_loss_prevention_discovery_config#aws_account_regex GoogleDataLossPreventionDiscoveryConfig#aws_account_regex}
	AwsAccountRegex *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexAwsAccountRegex `field:"optional" json:"awsAccountRegex" yaml:"awsAccountRegex"`
	// Regex to test the bucket name against. If empty, all buckets match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_data_loss_prevention_discovery_config#bucket_name_regex GoogleDataLossPreventionDiscoveryConfig#bucket_name_regex}
	BucketNameRegex *string `field:"optional" json:"bucketNameRegex" yaml:"bucketNameRegex"`
}

