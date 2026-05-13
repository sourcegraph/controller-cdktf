package googledialogflowgenerator


type GoogleDialogflowGeneratorSummarizationContext struct {
	// few_shot_examples block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_generator#few_shot_examples GoogleDialogflowGenerator#few_shot_examples}
	FewShotExamples interface{} `field:"optional" json:"fewShotExamples" yaml:"fewShotExamples"`
	// Optional.
	//
	// The target language of the generated summary. The language code for conversation will be used if this field is empty. Supported 2.0 and later versions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_generator#output_language_code GoogleDialogflowGenerator#output_language_code}
	OutputLanguageCode *string `field:"optional" json:"outputLanguageCode" yaml:"outputLanguageCode"`
	// summarization_sections block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_generator#summarization_sections GoogleDialogflowGenerator#summarization_sections}
	SummarizationSections interface{} `field:"optional" json:"summarizationSections" yaml:"summarizationSections"`
	// Optional. Version of the feature. If not set, default to latest version. Current candidates are ["1.0"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_generator#version GoogleDialogflowGenerator#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

