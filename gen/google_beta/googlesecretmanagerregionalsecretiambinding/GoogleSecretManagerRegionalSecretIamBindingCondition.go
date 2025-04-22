package googlesecretmanagerregionalsecretiambinding


type GoogleSecretManagerRegionalSecretIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_secret_manager_regional_secret_iam_binding#expression GoogleSecretManagerRegionalSecretIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_secret_manager_regional_secret_iam_binding#title GoogleSecretManagerRegionalSecretIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_secret_manager_regional_secret_iam_binding#description GoogleSecretManagerRegionalSecretIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

