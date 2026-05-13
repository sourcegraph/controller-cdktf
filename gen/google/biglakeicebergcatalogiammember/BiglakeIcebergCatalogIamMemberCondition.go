package biglakeicebergcatalogiammember


type BiglakeIcebergCatalogIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/biglake_iceberg_catalog_iam_member#expression BiglakeIcebergCatalogIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/biglake_iceberg_catalog_iam_member#title BiglakeIcebergCatalogIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/biglake_iceberg_catalog_iam_member#description BiglakeIcebergCatalogIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

