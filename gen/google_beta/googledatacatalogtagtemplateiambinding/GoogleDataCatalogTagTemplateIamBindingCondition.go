package googledatacatalogtagtemplateiambinding


type GoogleDataCatalogTagTemplateIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_catalog_tag_template_iam_binding#expression GoogleDataCatalogTagTemplateIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_catalog_tag_template_iam_binding#title GoogleDataCatalogTagTemplateIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_catalog_tag_template_iam_binding#description GoogleDataCatalogTagTemplateIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

