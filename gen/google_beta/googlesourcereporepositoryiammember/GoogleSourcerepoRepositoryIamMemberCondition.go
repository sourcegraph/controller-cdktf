package googlesourcereporepositoryiammember


type GoogleSourcerepoRepositoryIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sourcerepo_repository_iam_member#expression GoogleSourcerepoRepositoryIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sourcerepo_repository_iam_member#title GoogleSourcerepoRepositoryIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sourcerepo_repository_iam_member#description GoogleSourcerepoRepositoryIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

