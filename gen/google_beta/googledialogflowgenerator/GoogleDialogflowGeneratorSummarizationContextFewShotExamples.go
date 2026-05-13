package googledialogflowgenerator


type GoogleDialogflowGeneratorSummarizationContextFewShotExamples struct {
	// output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_generator#output GoogleDialogflowGenerator#output}
	Output *GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutput `field:"required" json:"output" yaml:"output"`
	// conversation_context block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_generator#conversation_context GoogleDialogflowGenerator#conversation_context}
	ConversationContext *GoogleDialogflowGeneratorSummarizationContextFewShotExamplesConversationContext `field:"optional" json:"conversationContext" yaml:"conversationContext"`
	// Optional.
	//
	// Key is the placeholder field name in input, value is the value of the placeholder. E.g. instruction contains "@price", and ingested data has <"price", "10">
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_generator#extra_info GoogleDialogflowGenerator#extra_info}
	ExtraInfo *map[string]*string `field:"optional" json:"extraInfo" yaml:"extraInfo"`
	// summarization_section_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_generator#summarization_section_list GoogleDialogflowGenerator#summarization_section_list}
	SummarizationSectionList *GoogleDialogflowGeneratorSummarizationContextFewShotExamplesSummarizationSectionListStruct `field:"optional" json:"summarizationSectionList" yaml:"summarizationSectionList"`
}

