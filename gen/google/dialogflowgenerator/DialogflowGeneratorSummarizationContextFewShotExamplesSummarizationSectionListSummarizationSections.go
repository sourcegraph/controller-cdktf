package dialogflowgenerator


type DialogflowGeneratorSummarizationContextFewShotExamplesSummarizationSectionListSummarizationSections struct {
	// Optional. Definition of the section, for example, "what the customer needs help with or has question about.".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_generator#definition DialogflowGenerator#definition}
	Definition *string `field:"optional" json:"definition" yaml:"definition"`
	// Optional. Name of the section, for example, "situation".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_generator#key DialogflowGenerator#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Optional. Type of the summarization section. Possible values: ["SITUATION", "ACTION", "RESOLUTION", "REASON_FOR_CANCELLATION", "CUSTOMER_SATISFACTION", "ENTITIES", "CUSTOMER_DEFINED", "SITUATION_CONCISE", "ACTION_CONCISE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_generator#type DialogflowGenerator#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

