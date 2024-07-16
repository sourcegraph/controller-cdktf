package googleiapwebbackendserviceiammember


type GoogleIapWebBackendServiceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_web_backend_service_iam_member#expression GoogleIapWebBackendServiceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_web_backend_service_iam_member#title GoogleIapWebBackendServiceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_web_backend_service_iam_member#description GoogleIapWebBackendServiceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

