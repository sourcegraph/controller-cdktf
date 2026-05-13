package googlecesexample


type GoogleCesExampleMessagesChunks struct {
	// agent_transfer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_example#agent_transfer GoogleCesExample#agent_transfer}
	AgentTransfer *GoogleCesExampleMessagesChunksAgentTransfer `field:"optional" json:"agentTransfer" yaml:"agentTransfer"`
	// image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_example#image GoogleCesExample#image}
	Image *GoogleCesExampleMessagesChunksImage `field:"optional" json:"image" yaml:"image"`
	// Text data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_example#text GoogleCesExample#text}
	Text *string `field:"optional" json:"text" yaml:"text"`
	// tool_call block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_example#tool_call GoogleCesExample#tool_call}
	ToolCall *GoogleCesExampleMessagesChunksToolCall `field:"optional" json:"toolCall" yaml:"toolCall"`
	// tool_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_example#tool_response GoogleCesExample#tool_response}
	ToolResponse *GoogleCesExampleMessagesChunksToolResponse `field:"optional" json:"toolResponse" yaml:"toolResponse"`
	// A struct represents variables that were updated in the conversation, keyed by variable names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_example#updated_variables GoogleCesExample#updated_variables}
	UpdatedVariables *string `field:"optional" json:"updatedVariables" yaml:"updatedVariables"`
}

