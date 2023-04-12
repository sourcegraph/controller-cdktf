package googlecomputebackendserviceiambinding


type GoogleComputeBackendServiceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_backend_service_iam_binding#expression GoogleComputeBackendServiceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_backend_service_iam_binding#title GoogleComputeBackendServiceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_backend_service_iam_binding#description GoogleComputeBackendServiceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

