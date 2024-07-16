package googlestoragebucket


type GoogleStorageBucketLogging struct {
	// The bucket that will receive log objects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_bucket#log_bucket GoogleStorageBucket#log_bucket}
	LogBucket *string `field:"required" json:"logBucket" yaml:"logBucket"`
	// The object prefix for log objects.
	//
	// If it's not provided, by default Google Cloud Storage sets this to this bucket's name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_bucket#log_object_prefix GoogleStorageBucket#log_object_prefix}
	LogObjectPrefix *string `field:"optional" json:"logObjectPrefix" yaml:"logObjectPrefix"`
}

