package googlecesguardrail


type GoogleCesGuardrailCodeCallback struct {
	// after_agent_callback block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_guardrail#after_agent_callback GoogleCesGuardrail#after_agent_callback}
	AfterAgentCallback *GoogleCesGuardrailCodeCallbackAfterAgentCallback `field:"optional" json:"afterAgentCallback" yaml:"afterAgentCallback"`
	// after_model_callback block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_guardrail#after_model_callback GoogleCesGuardrail#after_model_callback}
	AfterModelCallback *GoogleCesGuardrailCodeCallbackAfterModelCallback `field:"optional" json:"afterModelCallback" yaml:"afterModelCallback"`
	// before_agent_callback block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_guardrail#before_agent_callback GoogleCesGuardrail#before_agent_callback}
	BeforeAgentCallback *GoogleCesGuardrailCodeCallbackBeforeAgentCallback `field:"optional" json:"beforeAgentCallback" yaml:"beforeAgentCallback"`
	// before_model_callback block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_guardrail#before_model_callback GoogleCesGuardrail#before_model_callback}
	BeforeModelCallback *GoogleCesGuardrailCodeCallbackBeforeModelCallback `field:"optional" json:"beforeModelCallback" yaml:"beforeModelCallback"`
}

