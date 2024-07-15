package googleruntimeconfigconfigiambinding


type GoogleRuntimeconfigConfigIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_runtimeconfig_config_iam_binding#expression GoogleRuntimeconfigConfigIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_runtimeconfig_config_iam_binding#title GoogleRuntimeconfigConfigIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_runtimeconfig_config_iam_binding#description GoogleRuntimeconfigConfigIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

