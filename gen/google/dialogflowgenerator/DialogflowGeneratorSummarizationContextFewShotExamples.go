package dialogflowgenerator


type DialogflowGeneratorSummarizationContextFewShotExamples struct {
	// output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_generator#output DialogflowGenerator#output}
	Output *DialogflowGeneratorSummarizationContextFewShotExamplesOutput `field:"required" json:"output" yaml:"output"`
	// conversation_context block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_generator#conversation_context DialogflowGenerator#conversation_context}
	ConversationContext *DialogflowGeneratorSummarizationContextFewShotExamplesConversationContext `field:"optional" json:"conversationContext" yaml:"conversationContext"`
	// Optional.
	//
	// Key is the placeholder field name in input, value is the value of the placeholder. E.g. instruction contains "@price", and ingested data has <"price", "10">
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_generator#extra_info DialogflowGenerator#extra_info}
	ExtraInfo *map[string]*string `field:"optional" json:"extraInfo" yaml:"extraInfo"`
	// summarization_section_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_generator#summarization_section_list DialogflowGenerator#summarization_section_list}
	SummarizationSectionList *DialogflowGeneratorSummarizationContextFewShotExamplesSummarizationSectionListStruct `field:"optional" json:"summarizationSectionList" yaml:"summarizationSectionList"`
}

