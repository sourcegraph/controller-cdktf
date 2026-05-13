package googledialogflowcxflow


type GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesText struct {
	// A collection of text response variants.
	//
	// If multiple variants are defined, only one text response variant is returned at runtime.
	// required: true
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_flow#text GoogleDialogflowCxFlow#text}
	Text *[]*string `field:"optional" json:"text" yaml:"text"`
}

