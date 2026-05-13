package googlebiglakeicebergcatalogiambinding


type GoogleBiglakeIcebergCatalogIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_biglake_iceberg_catalog_iam_binding#expression GoogleBiglakeIcebergCatalogIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_biglake_iceberg_catalog_iam_binding#title GoogleBiglakeIcebergCatalogIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_biglake_iceberg_catalog_iam_binding#description GoogleBiglakeIcebergCatalogIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

