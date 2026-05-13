package googlebiglakeicebergtableiammember


type GoogleBiglakeIcebergTableIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_biglake_iceberg_table_iam_member#expression GoogleBiglakeIcebergTableIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_biglake_iceberg_table_iam_member#title GoogleBiglakeIcebergTableIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_biglake_iceberg_table_iam_member#description GoogleBiglakeIcebergTableIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

