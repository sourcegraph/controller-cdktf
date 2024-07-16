package googleprivatecacapooliambinding


type GooglePrivatecaCaPoolIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_ca_pool_iam_binding#expression GooglePrivatecaCaPoolIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_ca_pool_iam_binding#title GooglePrivatecaCaPoolIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_ca_pool_iam_binding#description GooglePrivatecaCaPoolIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

