package googlecesexample


type GoogleCesExampleMessagesChunksToolCallToolsetTool struct {
	// The resource name of the Toolset from which this tool is derived. Format: 'projects/{project}/locations/{location}/apps/{app}/toolsets/{toolset}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_example#toolset GoogleCesExample#toolset}
	Toolset *string `field:"required" json:"toolset" yaml:"toolset"`
	// The tool ID to filter the tools to retrieve the schema for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_example#tool_id GoogleCesExample#tool_id}
	ToolId *string `field:"optional" json:"toolId" yaml:"toolId"`
}

