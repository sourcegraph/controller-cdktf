package googlebeyondcorpapplicationiammember


type GoogleBeyondcorpApplicationIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_beyondcorp_application_iam_member#expression GoogleBeyondcorpApplicationIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_beyondcorp_application_iam_member#title GoogleBeyondcorpApplicationIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_beyondcorp_application_iam_member#description GoogleBeyondcorpApplicationIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

