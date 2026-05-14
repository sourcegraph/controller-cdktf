package googlecesevaluation


type GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunks struct {
	// agent_transfer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_ces_evaluation#agent_transfer GoogleCesEvaluation#agent_transfer}
	AgentTransfer *GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksAgentTransfer `field:"optional" json:"agentTransfer" yaml:"agentTransfer"`
	// blob block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_ces_evaluation#blob GoogleCesEvaluation#blob}
	Blob *GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksBlob `field:"optional" json:"blob" yaml:"blob"`
	// image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_ces_evaluation#image GoogleCesEvaluation#image}
	Image *GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksImage `field:"optional" json:"image" yaml:"image"`
	// Text data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_ces_evaluation#text GoogleCesEvaluation#text}
	Text *string `field:"optional" json:"text" yaml:"text"`
	// tool_call block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_ces_evaluation#tool_call GoogleCesEvaluation#tool_call}
	ToolCall *GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolCall `field:"optional" json:"toolCall" yaml:"toolCall"`
	// tool_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_ces_evaluation#tool_response GoogleCesEvaluation#tool_response}
	ToolResponse *GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponse `field:"optional" json:"toolResponse" yaml:"toolResponse"`
	// Updated variables in JSON object format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_ces_evaluation#updated_variables GoogleCesEvaluation#updated_variables}
	UpdatedVariables *map[string]*string `field:"optional" json:"updatedVariables" yaml:"updatedVariables"`
}

