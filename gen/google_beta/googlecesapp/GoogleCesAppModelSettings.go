package googlecesapp


type GoogleCesAppModelSettings struct {
	// The LLM model that the agent should use.
	//
	// If not set, the agent will inherit the model from its parent agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#model GoogleCesApp#model}
	Model *string `field:"optional" json:"model" yaml:"model"`
	// If set, this temperature will be used for the LLM model.
	//
	// Temperature
	// controls the randomness of the model's responses. Lower temperatures
	// produce responses that are more predictable. Higher temperatures produce
	// responses that are more creative.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#temperature GoogleCesApp#temperature}
	Temperature *float64 `field:"optional" json:"temperature" yaml:"temperature"`
}

