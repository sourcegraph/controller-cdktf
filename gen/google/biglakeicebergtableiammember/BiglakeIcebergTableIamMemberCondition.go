package biglakeicebergtableiammember


type BiglakeIcebergTableIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/biglake_iceberg_table_iam_member#expression BiglakeIcebergTableIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/biglake_iceberg_table_iam_member#title BiglakeIcebergTableIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/biglake_iceberg_table_iam_member#description BiglakeIcebergTableIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

