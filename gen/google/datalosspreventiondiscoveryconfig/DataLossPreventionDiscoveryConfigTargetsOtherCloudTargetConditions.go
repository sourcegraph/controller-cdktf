package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetConditions struct {
	// amazon_s3_bucket_conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/data_loss_prevention_discovery_config#amazon_s3_bucket_conditions DataLossPreventionDiscoveryConfig#amazon_s3_bucket_conditions}
	AmazonS3BucketConditions *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetConditionsAmazonS3BucketConditions `field:"optional" json:"amazonS3BucketConditions" yaml:"amazonS3BucketConditions"`
	// Duration format.
	//
	// Minimum age a resource must be before a profile can be generated. Value must be 1 hour or greater. Minimum age is not supported for Azure Blob Storage containers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/data_loss_prevention_discovery_config#min_age DataLossPreventionDiscoveryConfig#min_age}
	MinAge *string `field:"optional" json:"minAge" yaml:"minAge"`
}

