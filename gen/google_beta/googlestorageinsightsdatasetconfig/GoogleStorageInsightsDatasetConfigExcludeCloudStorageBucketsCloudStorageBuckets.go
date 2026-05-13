package googlestorageinsightsdatasetconfig


type GoogleStorageInsightsDatasetConfigExcludeCloudStorageBucketsCloudStorageBuckets struct {
	// The list of cloud storage bucket names to exclude in the DatasetConfig.
	//
	// Exactly one of the bucket_name and bucket_prefix_regex should be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_storage_insights_dataset_config#bucket_name GoogleStorageInsightsDatasetConfig#bucket_name}
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// The list of regex patterns for bucket names matching the regex.
	//
	// Regex should follow the syntax specified in google/re2 on GitHub.
	// Exactly one of the bucket_name and bucket_prefix_regex should be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_storage_insights_dataset_config#bucket_prefix_regex GoogleStorageInsightsDatasetConfig#bucket_prefix_regex}
	BucketPrefixRegex *string `field:"optional" json:"bucketPrefixRegex" yaml:"bucketPrefixRegex"`
}

