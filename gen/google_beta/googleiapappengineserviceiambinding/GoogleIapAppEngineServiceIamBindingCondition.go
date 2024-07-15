package googleiapappengineserviceiambinding


type GoogleIapAppEngineServiceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_app_engine_service_iam_binding#expression GoogleIapAppEngineServiceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_app_engine_service_iam_binding#title GoogleIapAppEngineServiceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_app_engine_service_iam_binding#description GoogleIapAppEngineServiceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

