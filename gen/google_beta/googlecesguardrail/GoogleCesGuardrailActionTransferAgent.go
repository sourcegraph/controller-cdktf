package googlecesguardrail


type GoogleCesGuardrailActionTransferAgent struct {
	// The name of the agent to transfer the conversation to.
	//
	// The agent must be
	// in the same app as the current agent.
	// Format:
	// 'projects/{project}/locations/{location}/apps/{app}/agents/{agent}'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_guardrail#agent GoogleCesGuardrail#agent}
	Agent *string `field:"required" json:"agent" yaml:"agent"`
}

