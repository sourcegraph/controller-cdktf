package googledialogflowcxpage


type GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesText struct {
	// A collection of text response variants.
	//
	// If multiple variants are defined, only one text response variant is returned at runtime.
	// required: true
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_page#text GoogleDialogflowCxPage#text}
	Text *[]*string `field:"optional" json:"text" yaml:"text"`
}

