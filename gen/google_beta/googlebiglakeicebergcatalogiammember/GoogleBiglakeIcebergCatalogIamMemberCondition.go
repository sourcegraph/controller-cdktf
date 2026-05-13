package googlebiglakeicebergcatalogiammember


type GoogleBiglakeIcebergCatalogIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_biglake_iceberg_catalog_iam_member#expression GoogleBiglakeIcebergCatalogIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_biglake_iceberg_catalog_iam_member#title GoogleBiglakeIcebergCatalogIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_biglake_iceberg_catalog_iam_member#description GoogleBiglakeIcebergCatalogIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

