package googlediscoveryenginechatengine


type GoogleDiscoveryEngineChatEngineChatEngineConfig struct {
	// agent_creation_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_discovery_engine_chat_engine#agent_creation_config GoogleDiscoveryEngineChatEngine#agent_creation_config}
	AgentCreationConfig *GoogleDiscoveryEngineChatEngineChatEngineConfigAgentCreationConfig `field:"optional" json:"agentCreationConfig" yaml:"agentCreationConfig"`
	// The resource name of an existing Dialogflow agent to link to this Chat Engine.
	//
	// Format: 'projects/<Project_ID>/locations/<Location_ID>/agents/<Agent_ID>'.
	// Exactly one of 'agent_creation_config' or 'dialogflow_agent_to_link' must be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_discovery_engine_chat_engine#dialogflow_agent_to_link GoogleDiscoveryEngineChatEngine#dialogflow_agent_to_link}
	DialogflowAgentToLink *string `field:"optional" json:"dialogflowAgentToLink" yaml:"dialogflowAgentToLink"`
}

