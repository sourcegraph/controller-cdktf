package dialogflowgenerator


type DialogflowGeneratorSummarizationContext struct {
	// few_shot_examples block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_generator#few_shot_examples DialogflowGenerator#few_shot_examples}
	FewShotExamples interface{} `field:"optional" json:"fewShotExamples" yaml:"fewShotExamples"`
	// Optional.
	//
	// The target language of the generated summary. The language code for conversation will be used if this field is empty. Supported 2.0 and later versions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_generator#output_language_code DialogflowGenerator#output_language_code}
	OutputLanguageCode *string `field:"optional" json:"outputLanguageCode" yaml:"outputLanguageCode"`
	// summarization_sections block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_generator#summarization_sections DialogflowGenerator#summarization_sections}
	SummarizationSections interface{} `field:"optional" json:"summarizationSections" yaml:"summarizationSections"`
	// Optional. Version of the feature. If not set, default to latest version. Current candidates are ["1.0"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_generator#version DialogflowGenerator#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

