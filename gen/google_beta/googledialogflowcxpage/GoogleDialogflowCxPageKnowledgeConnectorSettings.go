package googledialogflowcxpage


type GoogleDialogflowCxPageKnowledgeConnectorSettings struct {
	// data_store_connections block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_page#data_store_connections GoogleDialogflowCxPage#data_store_connections}
	DataStoreConnections interface{} `field:"optional" json:"dataStoreConnections" yaml:"dataStoreConnections"`
	// Whether Knowledge Connector is enabled or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_page#enabled GoogleDialogflowCxPage#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The target flow to transition to.
	//
	// Format: projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>.
	// This field is part of a union field 'target': Only one of 'targetPage' or 'targetFlow' may be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_page#target_flow GoogleDialogflowCxPage#target_flow}
	TargetFlow *string `field:"optional" json:"targetFlow" yaml:"targetFlow"`
	// The target page to transition to.
	//
	// Format: projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>/pages/<PageID>.
	// The page must be in the same host flow (the flow that owns this 'KnowledgeConnectorSettings').
	// This field is part of a union field 'target': Only one of 'targetPage' or 'targetFlow' may be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_page#target_page GoogleDialogflowCxPage#target_page}
	TargetPage *string `field:"optional" json:"targetPage" yaml:"targetPage"`
	// trigger_fulfillment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_page#trigger_fulfillment GoogleDialogflowCxPage#trigger_fulfillment}
	TriggerFulfillment *GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillment `field:"optional" json:"triggerFulfillment" yaml:"triggerFulfillment"`
}

