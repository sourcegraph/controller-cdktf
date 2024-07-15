package googledataplexaspecttypeiammember


type GoogleDataplexAspectTypeIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_aspect_type_iam_member#expression GoogleDataplexAspectTypeIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_aspect_type_iam_member#title GoogleDataplexAspectTypeIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_aspect_type_iam_member#description GoogleDataplexAspectTypeIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

