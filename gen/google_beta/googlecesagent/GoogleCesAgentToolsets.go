package googlecesagent


type GoogleCesAgentToolsets struct {
	// The resource name of the toolset. Format: 'projects/{project}/locations/{location}/apps/{app}/toolsets/{toolset}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_agent#toolset GoogleCesAgent#toolset}
	Toolset *string `field:"required" json:"toolset" yaml:"toolset"`
	// The tools IDs to filter the toolset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_agent#tool_ids GoogleCesAgent#tool_ids}
	ToolIds *[]*string `field:"optional" json:"toolIds" yaml:"toolIds"`
}

