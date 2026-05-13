package cesexample


type CesExampleMessagesChunksToolResponseToolsetTool struct {
	// The resource name of the Toolset from which this tool is derived. Format: 'projects/{project}/locations/{location}/apps/{app}/toolsets/{toolset}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_example#toolset CesExample#toolset}
	Toolset *string `field:"required" json:"toolset" yaml:"toolset"`
	// The tool ID to filter the tools to retrieve the schema for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_example#tool_id CesExample#tool_id}
	ToolId *string `field:"optional" json:"toolId" yaml:"toolId"`
}

