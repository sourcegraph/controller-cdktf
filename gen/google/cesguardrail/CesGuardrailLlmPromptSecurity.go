package cesguardrail


type CesGuardrailLlmPromptSecurity struct {
	// custom_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#custom_policy CesGuardrail#custom_policy}
	CustomPolicy *CesGuardrailLlmPromptSecurityCustomPolicy `field:"optional" json:"customPolicy" yaml:"customPolicy"`
	// default_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#default_settings CesGuardrail#default_settings}
	DefaultSettings *CesGuardrailLlmPromptSecurityDefaultSettings `field:"optional" json:"defaultSettings" yaml:"defaultSettings"`
}

