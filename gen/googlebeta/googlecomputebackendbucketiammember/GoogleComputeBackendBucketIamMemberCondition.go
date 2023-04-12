package googlecomputebackendbucketiammember


type GoogleComputeBackendBucketIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_backend_bucket_iam_member#expression GoogleComputeBackendBucketIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_backend_bucket_iam_member#title GoogleComputeBackendBucketIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_backend_bucket_iam_member#description GoogleComputeBackendBucketIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

