package googleiapwebtypeappengineiambinding


type GoogleIapWebTypeAppEngineIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_iap_web_type_app_engine_iam_binding#expression GoogleIapWebTypeAppEngineIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_iap_web_type_app_engine_iam_binding#title GoogleIapWebTypeAppEngineIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_iap_web_type_app_engine_iam_binding#description GoogleIapWebTypeAppEngineIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

