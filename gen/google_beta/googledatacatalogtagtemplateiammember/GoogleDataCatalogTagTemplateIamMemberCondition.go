package googledatacatalogtagtemplateiammember


type GoogleDataCatalogTagTemplateIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_catalog_tag_template_iam_member#expression GoogleDataCatalogTagTemplateIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_catalog_tag_template_iam_member#title GoogleDataCatalogTagTemplateIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_catalog_tag_template_iam_member#description GoogleDataCatalogTagTemplateIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

