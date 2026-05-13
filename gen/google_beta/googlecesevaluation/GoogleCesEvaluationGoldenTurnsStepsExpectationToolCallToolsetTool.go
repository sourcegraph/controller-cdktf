package googlecesevaluation


type GoogleCesEvaluationGoldenTurnsStepsExpectationToolCallToolsetTool struct {
	// The resource name of the Toolset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#toolset GoogleCesEvaluation#toolset}
	Toolset *string `field:"required" json:"toolset" yaml:"toolset"`
	// The tool ID to filter the tools to retrieve the schema for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#tool_id GoogleCesEvaluation#tool_id}
	ToolId *string `field:"optional" json:"toolId" yaml:"toolId"`
}

