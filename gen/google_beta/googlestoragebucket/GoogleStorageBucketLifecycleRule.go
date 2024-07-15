package googlestoragebucket


type GoogleStorageBucketLifecycleRule struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_bucket#action GoogleStorageBucket#action}
	Action *GoogleStorageBucketLifecycleRuleAction `field:"required" json:"action" yaml:"action"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_bucket#condition GoogleStorageBucket#condition}
	Condition *GoogleStorageBucketLifecycleRuleCondition `field:"required" json:"condition" yaml:"condition"`
}

