package googledataplexglossaryiambinding


type GoogleDataplexGlossaryIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_glossary_iam_binding#expression GoogleDataplexGlossaryIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_glossary_iam_binding#title GoogleDataplexGlossaryIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_glossary_iam_binding#description GoogleDataplexGlossaryIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

