package googledatacatalogtaxonomyiammember


type GoogleDataCatalogTaxonomyIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_catalog_taxonomy_iam_member#expression GoogleDataCatalogTaxonomyIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_catalog_taxonomy_iam_member#title GoogleDataCatalogTaxonomyIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_catalog_taxonomy_iam_member#description GoogleDataCatalogTaxonomyIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

