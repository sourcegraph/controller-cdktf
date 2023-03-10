package googleiapappengineversioniammember


type GoogleIapAppEngineVersionIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_iap_app_engine_version_iam_member#expression GoogleIapAppEngineVersionIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_iap_app_engine_version_iam_member#title GoogleIapAppEngineVersionIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_iap_app_engine_version_iam_member#description GoogleIapAppEngineVersionIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

