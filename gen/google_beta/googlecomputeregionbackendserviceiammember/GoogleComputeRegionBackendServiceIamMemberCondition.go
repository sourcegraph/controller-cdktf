package googlecomputeregionbackendserviceiammember


type GoogleComputeRegionBackendServiceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_backend_service_iam_member#expression GoogleComputeRegionBackendServiceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_backend_service_iam_member#title GoogleComputeRegionBackendServiceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_backend_service_iam_member#description GoogleComputeRegionBackendServiceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

