package googlecesguardrail


type GoogleCesGuardrailLlmPromptSecurity struct {
	// custom_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_guardrail#custom_policy GoogleCesGuardrail#custom_policy}
	CustomPolicy *GoogleCesGuardrailLlmPromptSecurityCustomPolicy `field:"optional" json:"customPolicy" yaml:"customPolicy"`
	// default_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_guardrail#default_settings GoogleCesGuardrail#default_settings}
	DefaultSettings *GoogleCesGuardrailLlmPromptSecurityDefaultSettings `field:"optional" json:"defaultSettings" yaml:"defaultSettings"`
}

