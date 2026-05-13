package googledialogflowgenerator


type GoogleDialogflowGeneratorSummarizationContextFewShotExamplesConversationContextMessageEntries struct {
	// Optional. Create time of the message entry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_generator#create_time GoogleDialogflowGenerator#create_time}
	CreateTime *string `field:"optional" json:"createTime" yaml:"createTime"`
	// Optional. The language of the text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_generator#language_code GoogleDialogflowGenerator#language_code}
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// Optional. Participant role of the message. Possible values: ["HUMAN_AGENT", "AUTOMATED_AGENT", "END_USER"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_generator#role GoogleDialogflowGenerator#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
	// Optional. Transcript content of the message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_generator#text GoogleDialogflowGenerator#text}
	Text *string `field:"optional" json:"text" yaml:"text"`
}

