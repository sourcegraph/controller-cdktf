package googleservicedirectoryserviceiambinding


type GoogleServiceDirectoryServiceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_service_directory_service_iam_binding#expression GoogleServiceDirectoryServiceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_service_directory_service_iam_binding#title GoogleServiceDirectoryServiceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_service_directory_service_iam_binding#description GoogleServiceDirectoryServiceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

