package googlecomputebackendbucketiambinding


type GoogleComputeBackendBucketIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_backend_bucket_iam_binding#expression GoogleComputeBackendBucketIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_backend_bucket_iam_binding#title GoogleComputeBackendBucketIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_backend_bucket_iam_binding#description GoogleComputeBackendBucketIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

