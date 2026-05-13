package googleiaplocationwebiambinding


type GoogleIapLocationWebIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iap_location_web_iam_binding#expression GoogleIapLocationWebIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iap_location_web_iam_binding#title GoogleIapLocationWebIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iap_location_web_iam_binding#description GoogleIapLocationWebIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

