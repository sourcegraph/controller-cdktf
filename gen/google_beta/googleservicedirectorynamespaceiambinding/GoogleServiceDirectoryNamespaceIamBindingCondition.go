package googleservicedirectorynamespaceiambinding


type GoogleServiceDirectoryNamespaceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_service_directory_namespace_iam_binding#expression GoogleServiceDirectoryNamespaceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_service_directory_namespace_iam_binding#title GoogleServiceDirectoryNamespaceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_service_directory_namespace_iam_binding#description GoogleServiceDirectoryNamespaceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

