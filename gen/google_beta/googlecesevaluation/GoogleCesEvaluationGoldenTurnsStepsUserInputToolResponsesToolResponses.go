package googlecesevaluation


type GoogleCesEvaluationGoldenTurnsStepsUserInputToolResponsesToolResponses struct {
	// Required. The tool execution result in JSON.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#response GoogleCesEvaluation#response}
	Response *map[string]*string `field:"required" json:"response" yaml:"response"`
	// Optional. Matching ID of the tool call.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#id GoogleCesEvaluation#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Name of the tool to execute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#tool GoogleCesEvaluation#tool}
	Tool *string `field:"optional" json:"tool" yaml:"tool"`
	// toolset_tool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#toolset_tool GoogleCesEvaluation#toolset_tool}
	ToolsetTool *GoogleCesEvaluationGoldenTurnsStepsUserInputToolResponsesToolResponsesToolsetTool `field:"optional" json:"toolsetTool" yaml:"toolsetTool"`
}

