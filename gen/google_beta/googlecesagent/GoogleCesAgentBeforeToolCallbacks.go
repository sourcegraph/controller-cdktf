package googlecesagent


type GoogleCesAgentBeforeToolCallbacks struct {
	// The python code to execute for the callback.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_agent#python_code GoogleCesAgent#python_code}
	PythonCode *string `field:"required" json:"pythonCode" yaml:"pythonCode"`
	// Human-readable description of the callback.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_agent#description GoogleCesAgent#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the callback is disabled. Disabled callbacks are ignored by the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_agent#disabled GoogleCesAgent#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
}

