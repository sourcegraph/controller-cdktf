package googledatacatalogpolicytagiammember


type GoogleDataCatalogPolicyTagIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_data_catalog_policy_tag_iam_member#expression GoogleDataCatalogPolicyTagIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_data_catalog_policy_tag_iam_member#title GoogleDataCatalogPolicyTagIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_data_catalog_policy_tag_iam_member#description GoogleDataCatalogPolicyTagIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

