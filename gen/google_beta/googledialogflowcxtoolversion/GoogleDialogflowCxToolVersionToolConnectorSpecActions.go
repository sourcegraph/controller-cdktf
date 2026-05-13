package googledialogflowcxtoolversion


type GoogleDialogflowCxToolVersionToolConnectorSpecActions struct {
	// ID of a Connection action for the tool to use.
	//
	// This field is part of a required union field 'action_spec'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool_version#connection_action_id GoogleDialogflowCxToolVersion#connection_action_id}
	ConnectionActionId *string `field:"optional" json:"connectionActionId" yaml:"connectionActionId"`
	// entity_operation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool_version#entity_operation GoogleDialogflowCxToolVersion#entity_operation}
	EntityOperation *GoogleDialogflowCxToolVersionToolConnectorSpecActionsEntityOperation `field:"optional" json:"entityOperation" yaml:"entityOperation"`
	// Entity fields to use as inputs for the operation.
	//
	// If no fields are specified, all fields of the Entity will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool_version#input_fields GoogleDialogflowCxToolVersion#input_fields}
	InputFields *[]*string `field:"optional" json:"inputFields" yaml:"inputFields"`
	// Entity fields to return from the operation. If no fields are specified, all fields of the Entity will be returned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool_version#output_fields GoogleDialogflowCxToolVersion#output_fields}
	OutputFields *[]*string `field:"optional" json:"outputFields" yaml:"outputFields"`
}

