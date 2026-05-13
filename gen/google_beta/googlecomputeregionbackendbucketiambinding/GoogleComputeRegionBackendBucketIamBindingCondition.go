package googlecomputeregionbackendbucketiambinding


type GoogleComputeRegionBackendBucketIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_region_backend_bucket_iam_binding#expression GoogleComputeRegionBackendBucketIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_region_backend_bucket_iam_binding#title GoogleComputeRegionBackendBucketIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_region_backend_bucket_iam_binding#description GoogleComputeRegionBackendBucketIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

