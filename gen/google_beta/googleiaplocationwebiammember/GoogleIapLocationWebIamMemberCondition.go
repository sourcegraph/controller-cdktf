package googleiaplocationwebiammember


type GoogleIapLocationWebIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iap_location_web_iam_member#expression GoogleIapLocationWebIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iap_location_web_iam_member#title GoogleIapLocationWebIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iap_location_web_iam_member#description GoogleIapLocationWebIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

