package googledialogflowcxflow


type GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentSetParameterActions struct {
	// Display name of the parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_flow#parameter GoogleDialogflowCxFlow#parameter}
	Parameter *string `field:"optional" json:"parameter" yaml:"parameter"`
	// The new JSON-encoded value of the parameter. A null value clears the parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_flow#value GoogleDialogflowCxFlow#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

