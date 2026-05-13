package googlecesguardrail


type GoogleCesGuardrailActionRespondImmediatelyResponses struct {
	// Text for the agent to respond with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_guardrail#text GoogleCesGuardrail#text}
	Text *string `field:"required" json:"text" yaml:"text"`
	// Whether the response is disabled. Disabled responses are not used by the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_guardrail#disabled GoogleCesGuardrail#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
}

