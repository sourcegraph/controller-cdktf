package googlestoragebucket


type GoogleStorageBucketLifecycleRule struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_bucket#action GoogleStorageBucket#action}
	Action *GoogleStorageBucketLifecycleRuleAction `field:"required" json:"action" yaml:"action"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_bucket#condition GoogleStorageBucket#condition}
	Condition *GoogleStorageBucketLifecycleRuleCondition `field:"required" json:"condition" yaml:"condition"`
}

