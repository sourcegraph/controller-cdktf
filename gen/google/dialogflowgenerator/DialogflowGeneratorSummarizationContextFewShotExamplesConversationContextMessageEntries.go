package dialogflowgenerator


type DialogflowGeneratorSummarizationContextFewShotExamplesConversationContextMessageEntries struct {
	// Optional. Create time of the message entry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_generator#create_time DialogflowGenerator#create_time}
	CreateTime *string `field:"optional" json:"createTime" yaml:"createTime"`
	// Optional. The language of the text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_generator#language_code DialogflowGenerator#language_code}
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// Optional. Participant role of the message. Possible values: ["HUMAN_AGENT", "AUTOMATED_AGENT", "END_USER"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_generator#role DialogflowGenerator#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
	// Optional. Transcript content of the message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_generator#text DialogflowGenerator#text}
	Text *string `field:"optional" json:"text" yaml:"text"`
}

