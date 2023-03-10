package googleiapappengineserviceiambinding


type GoogleIapAppEngineServiceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_iap_app_engine_service_iam_binding#expression GoogleIapAppEngineServiceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_iap_app_engine_service_iam_binding#title GoogleIapAppEngineServiceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_iap_app_engine_service_iam_binding#description GoogleIapAppEngineServiceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

