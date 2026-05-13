package cesexample


type CesExampleMessagesChunksToolCall struct {
	// The input parameters and values for the tool in JSON object format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_example#args CesExample#args}
	Args *string `field:"optional" json:"args" yaml:"args"`
	// The unique identifier of the tool call.
	//
	// If populated, the client should
	// return the execution result with the matching ID in
	// ToolResponse.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_example#id CesExample#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of the tool to execute. Format: 'projects/{project}/locations/{location}/apps/{app}/tools/{tool}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_example#tool CesExample#tool}
	Tool *string `field:"optional" json:"tool" yaml:"tool"`
	// toolset_tool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_example#toolset_tool CesExample#toolset_tool}
	ToolsetTool *CesExampleMessagesChunksToolCallToolsetTool `field:"optional" json:"toolsetTool" yaml:"toolsetTool"`
}

