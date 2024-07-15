package googledataplexaspecttypeiambinding


type GoogleDataplexAspectTypeIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_aspect_type_iam_binding#expression GoogleDataplexAspectTypeIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_aspect_type_iam_binding#title GoogleDataplexAspectTypeIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_aspect_type_iam_binding#description GoogleDataplexAspectTypeIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

