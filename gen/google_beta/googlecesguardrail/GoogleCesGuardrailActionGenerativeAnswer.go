package googlecesguardrail


type GoogleCesGuardrailActionGenerativeAnswer struct {
	// The prompt to use for the generative answer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_guardrail#prompt GoogleCesGuardrail#prompt}
	Prompt *string `field:"required" json:"prompt" yaml:"prompt"`
}

