package googlesourcereporepositoryiambinding


type GoogleSourcerepoRepositoryIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sourcerepo_repository_iam_binding#expression GoogleSourcerepoRepositoryIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sourcerepo_repository_iam_binding#title GoogleSourcerepoRepositoryIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sourcerepo_repository_iam_binding#description GoogleSourcerepoRepositoryIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

