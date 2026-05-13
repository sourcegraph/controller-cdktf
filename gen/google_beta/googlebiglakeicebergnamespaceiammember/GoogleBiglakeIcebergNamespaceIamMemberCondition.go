package googlebiglakeicebergnamespaceiammember


type GoogleBiglakeIcebergNamespaceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_biglake_iceberg_namespace_iam_member#expression GoogleBiglakeIcebergNamespaceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_biglake_iceberg_namespace_iam_member#title GoogleBiglakeIcebergNamespaceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_biglake_iceberg_namespace_iam_member#description GoogleBiglakeIcebergNamespaceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

