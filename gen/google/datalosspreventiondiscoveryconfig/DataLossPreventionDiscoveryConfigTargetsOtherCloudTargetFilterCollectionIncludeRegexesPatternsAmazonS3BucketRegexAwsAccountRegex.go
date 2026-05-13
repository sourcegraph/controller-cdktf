package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionIncludeRegexesPatternsAmazonS3BucketRegexAwsAccountRegex struct {
	// Regex to test the AWS account ID against. If empty, all accounts match. Example: arn:aws:organizations::123:account/o-b2c3d4/345.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/data_loss_prevention_discovery_config#account_id_regex DataLossPreventionDiscoveryConfig#account_id_regex}
	AccountIdRegex *string `field:"optional" json:"accountIdRegex" yaml:"accountIdRegex"`
}

