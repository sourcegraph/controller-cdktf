package servicedirectorynamespaceiammember


type ServiceDirectoryNamespaceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/service_directory_namespace_iam_member#expression ServiceDirectoryNamespaceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/service_directory_namespace_iam_member#title ServiceDirectoryNamespaceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/service_directory_namespace_iam_member#description ServiceDirectoryNamespaceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

