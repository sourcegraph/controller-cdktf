package googlecomputebackendserviceiammember


type GoogleComputeBackendServiceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_backend_service_iam_member#expression GoogleComputeBackendServiceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_backend_service_iam_member#title GoogleComputeBackendServiceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_backend_service_iam_member#description GoogleComputeBackendServiceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

