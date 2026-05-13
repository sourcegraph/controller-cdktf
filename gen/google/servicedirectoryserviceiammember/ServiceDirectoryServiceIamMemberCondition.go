package servicedirectoryserviceiammember


type ServiceDirectoryServiceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/service_directory_service_iam_member#expression ServiceDirectoryServiceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/service_directory_service_iam_member#title ServiceDirectoryServiceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/service_directory_service_iam_member#description ServiceDirectoryServiceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

