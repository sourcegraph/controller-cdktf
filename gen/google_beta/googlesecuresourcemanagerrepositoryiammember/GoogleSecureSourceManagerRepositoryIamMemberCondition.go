package googlesecuresourcemanagerrepositoryiammember


type GoogleSecureSourceManagerRepositoryIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_secure_source_manager_repository_iam_member#expression GoogleSecureSourceManagerRepositoryIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_secure_source_manager_repository_iam_member#title GoogleSecureSourceManagerRepositoryIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_secure_source_manager_repository_iam_member#description GoogleSecureSourceManagerRepositoryIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

