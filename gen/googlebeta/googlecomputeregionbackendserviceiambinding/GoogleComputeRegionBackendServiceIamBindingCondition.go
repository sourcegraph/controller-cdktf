package googlecomputeregionbackendserviceiambinding


type GoogleComputeRegionBackendServiceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_region_backend_service_iam_binding#expression GoogleComputeRegionBackendServiceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_region_backend_service_iam_binding#title GoogleComputeRegionBackendServiceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_region_backend_service_iam_binding#description GoogleComputeRegionBackendServiceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

