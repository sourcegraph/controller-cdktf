package cesguardrail


type CesGuardrailActionGenerativeAnswer struct {
	// The prompt to use for the generative answer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#prompt CesGuardrail#prompt}
	Prompt *string `field:"required" json:"prompt" yaml:"prompt"`
}

