package googlebiglakeicebergnamespaceiambinding


type GoogleBiglakeIcebergNamespaceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_biglake_iceberg_namespace_iam_binding#expression GoogleBiglakeIcebergNamespaceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_biglake_iceberg_namespace_iam_binding#title GoogleBiglakeIcebergNamespaceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_biglake_iceberg_namespace_iam_binding#description GoogleBiglakeIcebergNamespaceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

