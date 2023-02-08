package googlestoragebucketiammember


type GoogleStorageBucketIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_bucket_iam_member#expression GoogleStorageBucketIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_bucket_iam_member#title GoogleStorageBucketIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_bucket_iam_member#description GoogleStorageBucketIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

