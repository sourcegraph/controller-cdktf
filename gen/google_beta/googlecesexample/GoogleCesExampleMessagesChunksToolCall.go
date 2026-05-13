package googlecesexample


type GoogleCesExampleMessagesChunksToolCall struct {
	// The input parameters and values for the tool in JSON object format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_example#args GoogleCesExample#args}
	Args *string `field:"optional" json:"args" yaml:"args"`
	// The unique identifier of the tool call.
	//
	// If populated, the client should
	// return the execution result with the matching ID in
	// ToolResponse.
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
	ToolsetTool *GoogleCesExampleMessagesChunksToolCallToolsetTool `field:"optional" json:"toolsetTool" yaml:"toolsetTool"`
}

