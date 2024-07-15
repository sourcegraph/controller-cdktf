package googledatacatalogentrygroupiambinding


type GoogleDataCatalogEntryGroupIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_catalog_entry_group_iam_binding#expression GoogleDataCatalogEntryGroupIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_catalog_entry_group_iam_binding#title GoogleDataCatalogEntryGroupIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_catalog_entry_group_iam_binding#description GoogleDataCatalogEntryGroupIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

