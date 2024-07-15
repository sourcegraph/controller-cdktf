package googledatacatalogtaxonomyiambinding


type GoogleDataCatalogTaxonomyIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_catalog_taxonomy_iam_binding#expression GoogleDataCatalogTaxonomyIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_catalog_taxonomy_iam_binding#title GoogleDataCatalogTaxonomyIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_catalog_taxonomy_iam_binding#description GoogleDataCatalogTaxonomyIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

