package googledataplexglossaryiammember


type GoogleDataplexGlossaryIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_glossary_iam_member#expression GoogleDataplexGlossaryIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_glossary_iam_member#title GoogleDataplexGlossaryIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_glossary_iam_member#description GoogleDataplexGlossaryIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

