package googledialogflowgenerator


type GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySections struct {
	// Required. Name of the section.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_generator#section GoogleDialogflowGenerator#section}
	Section *string `field:"required" json:"section" yaml:"section"`
	// Required. Summary text for the section.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_generator#summary GoogleDialogflowGenerator#summary}
	Summary *string `field:"required" json:"summary" yaml:"summary"`
}

