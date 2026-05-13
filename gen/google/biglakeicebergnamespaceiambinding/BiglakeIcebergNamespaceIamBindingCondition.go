package biglakeicebergnamespaceiambinding


type BiglakeIcebergNamespaceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/biglake_iceberg_namespace_iam_binding#expression BiglakeIcebergNamespaceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/biglake_iceberg_namespace_iam_binding#title BiglakeIcebergNamespaceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/biglake_iceberg_namespace_iam_binding#description BiglakeIcebergNamespaceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

