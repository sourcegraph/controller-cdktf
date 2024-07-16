package googledialogflowcxpage


type GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesOutputAudioText struct {
	// The SSML text to be synthesized. For more information, see SSML.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_page#ssml GoogleDialogflowCxPage#ssml}
	Ssml *string `field:"optional" json:"ssml" yaml:"ssml"`
	// The raw text to be synthesized.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_page#text GoogleDialogflowCxPage#text}
	Text *string `field:"optional" json:"text" yaml:"text"`
}

