package cesexample


type CesExampleMessagesChunks struct {
	// agent_transfer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_example#agent_transfer CesExample#agent_transfer}
	AgentTransfer *CesExampleMessagesChunksAgentTransfer `field:"optional" json:"agentTransfer" yaml:"agentTransfer"`
	// image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_example#image CesExample#image}
	Image *CesExampleMessagesChunksImage `field:"optional" json:"image" yaml:"image"`
	// Text data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_example#text CesExample#text}
	Text *string `field:"optional" json:"text" yaml:"text"`
	// tool_call block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_example#tool_call CesExample#tool_call}
	ToolCall *CesExampleMessagesChunksToolCall `field:"optional" json:"toolCall" yaml:"toolCall"`
	// tool_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_example#tool_response CesExample#tool_response}
	ToolResponse *CesExampleMessagesChunksToolResponse `field:"optional" json:"toolResponse" yaml:"toolResponse"`
	// A struct represents variables that were updated in the conversation, keyed by variable names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_example#updated_variables CesExample#updated_variables}
	UpdatedVariables *string `field:"optional" json:"updatedVariables" yaml:"updatedVariables"`
}

