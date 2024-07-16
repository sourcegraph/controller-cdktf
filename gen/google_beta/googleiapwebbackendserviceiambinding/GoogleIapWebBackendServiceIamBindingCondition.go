package googleiapwebbackendserviceiambinding


type GoogleIapWebBackendServiceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_web_backend_service_iam_binding#expression GoogleIapWebBackendServiceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_web_backend_service_iam_binding#title GoogleIapWebBackendServiceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_web_backend_service_iam_binding#description GoogleIapWebBackendServiceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

