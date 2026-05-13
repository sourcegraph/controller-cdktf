package googledialogflowcxpage


type GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputAudioText struct {
	// The SSML text to be synthesized.
	//
	// For more information, see SSML.
	// This field is part of a union field 'source': Only one of 'text' or 'ssml' may be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_page#ssml GoogleDialogflowCxPage#ssml}
	Ssml *string `field:"optional" json:"ssml" yaml:"ssml"`
	// The raw text to be synthesized.
	//
	// This field is part of a union field 'source': Only one of 'text' or 'ssml' may be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_page#text GoogleDialogflowCxPage#text}
	Text *string `field:"optional" json:"text" yaml:"text"`
}

