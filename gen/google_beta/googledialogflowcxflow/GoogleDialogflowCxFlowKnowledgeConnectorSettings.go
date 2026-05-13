package googledialogflowcxflow


type GoogleDialogflowCxFlowKnowledgeConnectorSettings struct {
	// data_store_connections block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_flow#data_store_connections GoogleDialogflowCxFlow#data_store_connections}
	DataStoreConnections interface{} `field:"optional" json:"dataStoreConnections" yaml:"dataStoreConnections"`
	// Whether Knowledge Connector is enabled or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_flow#enabled GoogleDialogflowCxFlow#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The target flow to transition to.
	//
	// Format: projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>.
	// This field is part of a union field 'target': Only one of 'targetPage' or 'targetFlow' may be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_flow#target_flow GoogleDialogflowCxFlow#target_flow}
	TargetFlow *string `field:"optional" json:"targetFlow" yaml:"targetFlow"`
	// The target page to transition to.
	//
	// Format: projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>/pages/<PageID>.
	// The page must be in the same host flow (the flow that owns this 'KnowledgeConnectorSettings').
	// This field is part of a union field 'target': Only one of 'targetPage' or 'targetFlow' may be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_flow#target_page GoogleDialogflowCxFlow#target_page}
	TargetPage *string `field:"optional" json:"targetPage" yaml:"targetPage"`
	// trigger_fulfillment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_flow#trigger_fulfillment GoogleDialogflowCxFlow#trigger_fulfillment}
	TriggerFulfillment *GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillment `field:"optional" json:"triggerFulfillment" yaml:"triggerFulfillment"`
}

