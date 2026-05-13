package cesagent


type CesAgentRemoteDialogflowAgent struct {
	// The [Dialogflow](https://cloud.google.com/dialogflow/cx/docs/concept/console-conversational-agents agent resource name. Format: 'projects/{project}/locations/{location}/agents/{agent}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_agent#agent CesAgent#agent}
	Agent *string `field:"required" json:"agent" yaml:"agent"`
	// The flow ID of the flow in the Dialogflow agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_agent#flow_id CesAgent#flow_id}
	FlowId *string `field:"required" json:"flowId" yaml:"flowId"`
	// The environment ID of the Dialogflow agent be used for the agent execution.
	//
	// If not specified, the draft environment will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_agent#environment_id CesAgent#environment_id}
	EnvironmentId *string `field:"optional" json:"environmentId" yaml:"environmentId"`
	// The mapping of the app variables names to the Dialogflow session parameters names to be sent to the Dialogflow agent as input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_agent#input_variable_mapping CesAgent#input_variable_mapping}
	InputVariableMapping *map[string]*string `field:"optional" json:"inputVariableMapping" yaml:"inputVariableMapping"`
	// The mapping of the Dialogflow session parameters names to the app variables names to be sent back to the CES agent after the Dialogflow agent execution ends.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_agent#output_variable_mapping CesAgent#output_variable_mapping}
	OutputVariableMapping *map[string]*string `field:"optional" json:"outputVariableMapping" yaml:"outputVariableMapping"`
	// Indicates whether to respect the message-level interruption settings configured in the Dialogflow agent.
	//
	// * If false: all response messages from the Dialogflow agent follow the app-level barge-in settings. * If true: only response messages with ['allow_playback_interruption'](https://docs.cloud.google.com/dialogflow/cx/docs/reference/rpc/google.cloud.dialogflow.cx.v3#text) set to true will be interruptable, all other messages follow the app-level barge-in settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_agent#respect_response_interruption_settings CesAgent#respect_response_interruption_settings}
	RespectResponseInterruptionSettings interface{} `field:"optional" json:"respectResponseInterruptionSettings" yaml:"respectResponseInterruptionSettings"`
}

