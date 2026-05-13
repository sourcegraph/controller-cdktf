package cesguardrail


type CesGuardrailActionTransferAgent struct {
	// The name of the agent to transfer the conversation to.
	//
	// The agent must be
	// in the same app as the current agent.
	// Format:
	// 'projects/{project}/locations/{location}/apps/{app}/agents/{agent}'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#agent CesGuardrail#agent}
	Agent *string `field:"required" json:"agent" yaml:"agent"`
}

