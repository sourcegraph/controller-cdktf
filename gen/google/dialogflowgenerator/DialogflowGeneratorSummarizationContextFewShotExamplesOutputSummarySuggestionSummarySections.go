package dialogflowgenerator


type DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySections struct {
	// Required. Name of the section.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_generator#section DialogflowGenerator#section}
	Section *string `field:"required" json:"section" yaml:"section"`
	// Required. Summary text for the section.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_generator#summary DialogflowGenerator#summary}
	Summary *string `field:"required" json:"summary" yaml:"summary"`
}

