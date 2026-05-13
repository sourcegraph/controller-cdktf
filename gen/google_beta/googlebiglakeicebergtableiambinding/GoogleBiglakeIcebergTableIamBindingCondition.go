package googlebiglakeicebergtableiambinding


type GoogleBiglakeIcebergTableIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_biglake_iceberg_table_iam_binding#expression GoogleBiglakeIcebergTableIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_biglake_iceberg_table_iam_binding#title GoogleBiglakeIcebergTableIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_biglake_iceberg_table_iam_binding#description GoogleBiglakeIcebergTableIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

