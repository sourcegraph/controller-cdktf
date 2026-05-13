package googlecesexample


type GoogleCesExampleMessagesChunksToolResponse struct {
	// The tool execution result in JSON object format.
	//
	// Use "output" key to specify tool response and "error" key to specify
	// error details (if any). If "output" and "error" keys are not specified,
	// then whole "response" is treated as tool execution result.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_example#response GoogleCesExample#response}
	Response *string `field:"required" json:"response" yaml:"response"`
	// The matching ID of the tool call the response is for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_example#id GoogleCesExample#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of the tool to execute. Format: 'projects/{project}/locations/{location}/apps/{app}/tools/{tool}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_example#tool GoogleCesExample#tool}
	Tool *string `field:"optional" json:"tool" yaml:"tool"`
	// toolset_tool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_example#toolset_tool GoogleCesExample#toolset_tool}
	ToolsetTool *GoogleCesExampleMessagesChunksToolResponseToolsetTool `field:"optional" json:"toolsetTool" yaml:"toolsetTool"`
}

