package googleiapwebregionbackendserviceiammember


type GoogleIapWebRegionBackendServiceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_web_region_backend_service_iam_member#expression GoogleIapWebRegionBackendServiceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_web_region_backend_service_iam_member#title GoogleIapWebRegionBackendServiceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_web_region_backend_service_iam_member#description GoogleIapWebRegionBackendServiceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

