package googledatacatalogentrygroupiammember


type GoogleDataCatalogEntryGroupIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_data_catalog_entry_group_iam_member#expression GoogleDataCatalogEntryGroupIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_data_catalog_entry_group_iam_member#title GoogleDataCatalogEntryGroupIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_data_catalog_entry_group_iam_member#description GoogleDataCatalogEntryGroupIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

