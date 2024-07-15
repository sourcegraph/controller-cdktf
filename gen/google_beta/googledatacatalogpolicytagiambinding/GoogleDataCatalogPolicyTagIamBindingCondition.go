package googledatacatalogpolicytagiambinding


type GoogleDataCatalogPolicyTagIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_catalog_policy_tag_iam_binding#expression GoogleDataCatalogPolicyTagIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_catalog_policy_tag_iam_binding#title GoogleDataCatalogPolicyTagIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_catalog_policy_tag_iam_binding#description GoogleDataCatalogPolicyTagIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

