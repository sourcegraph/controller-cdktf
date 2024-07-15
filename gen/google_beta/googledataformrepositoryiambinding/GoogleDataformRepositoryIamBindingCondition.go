package googledataformrepositoryiambinding


type GoogleDataformRepositoryIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository_iam_binding#expression GoogleDataformRepositoryIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository_iam_binding#title GoogleDataformRepositoryIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository_iam_binding#description GoogleDataformRepositoryIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

