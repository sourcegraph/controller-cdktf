package cesguardrail


type CesGuardrailCodeCallbackBeforeAgentCallback struct {
	// The python code to execute for the callback.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#python_code CesGuardrail#python_code}
	PythonCode *string `field:"required" json:"pythonCode" yaml:"pythonCode"`
	// Human-readable description of the callback.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#description CesGuardrail#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the callback is disabled. Disabled callbacks are ignored by the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#disabled CesGuardrail#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
}

