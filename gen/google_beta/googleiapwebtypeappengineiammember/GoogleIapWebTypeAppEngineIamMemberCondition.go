package googleiapwebtypeappengineiammember


type GoogleIapWebTypeAppEngineIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_web_type_app_engine_iam_member#expression GoogleIapWebTypeAppEngineIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_web_type_app_engine_iam_member#title GoogleIapWebTypeAppEngineIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_web_type_app_engine_iam_member#description GoogleIapWebTypeAppEngineIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

