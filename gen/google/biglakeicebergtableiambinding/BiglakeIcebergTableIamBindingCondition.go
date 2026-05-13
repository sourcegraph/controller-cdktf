package biglakeicebergtableiambinding


type BiglakeIcebergTableIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/biglake_iceberg_table_iam_binding#expression BiglakeIcebergTableIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/biglake_iceberg_table_iam_binding#title BiglakeIcebergTableIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/biglake_iceberg_table_iam_binding#description BiglakeIcebergTableIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

