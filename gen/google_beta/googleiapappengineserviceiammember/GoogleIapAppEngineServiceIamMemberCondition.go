package googleiapappengineserviceiammember


type GoogleIapAppEngineServiceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_app_engine_service_iam_member#expression GoogleIapAppEngineServiceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_app_engine_service_iam_member#title GoogleIapAppEngineServiceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_app_engine_service_iam_member#description GoogleIapAppEngineServiceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

