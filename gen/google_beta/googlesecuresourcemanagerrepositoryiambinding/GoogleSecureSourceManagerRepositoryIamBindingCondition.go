package googlesecuresourcemanagerrepositoryiambinding


type GoogleSecureSourceManagerRepositoryIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_secure_source_manager_repository_iam_binding#expression GoogleSecureSourceManagerRepositoryIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_secure_source_manager_repository_iam_binding#title GoogleSecureSourceManagerRepositoryIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_secure_source_manager_repository_iam_binding#description GoogleSecureSourceManagerRepositoryIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

