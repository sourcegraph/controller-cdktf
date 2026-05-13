package googlecesevaluation


type GoogleCesEvaluationGoldenTurnsStepsExpectationToolCall struct {
	// The input parameters and values for the tool in JSON object format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#args GoogleCesEvaluation#args}
	Args *map[string]*string `field:"optional" json:"args" yaml:"args"`
	// The unique identifier of the tool call.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#id GoogleCesEvaluation#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The resource name of the tool. Note: 'tool' and 'toolset_tool' are mutually exclusive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#tool GoogleCesEvaluation#tool}
	Tool *string `field:"optional" json:"tool" yaml:"tool"`
	// toolset_tool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#toolset_tool GoogleCesEvaluation#toolset_tool}
	ToolsetTool *GoogleCesEvaluationGoldenTurnsStepsExpectationToolCallToolsetTool `field:"optional" json:"toolsetTool" yaml:"toolsetTool"`
}

