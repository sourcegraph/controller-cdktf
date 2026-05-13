package cesguardrail


type CesGuardrailCodeCallback struct {
	// after_agent_callback block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#after_agent_callback CesGuardrail#after_agent_callback}
	AfterAgentCallback *CesGuardrailCodeCallbackAfterAgentCallback `field:"optional" json:"afterAgentCallback" yaml:"afterAgentCallback"`
	// after_model_callback block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#after_model_callback CesGuardrail#after_model_callback}
	AfterModelCallback *CesGuardrailCodeCallbackAfterModelCallback `field:"optional" json:"afterModelCallback" yaml:"afterModelCallback"`
	// before_agent_callback block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#before_agent_callback CesGuardrail#before_agent_callback}
	BeforeAgentCallback *CesGuardrailCodeCallbackBeforeAgentCallback `field:"optional" json:"beforeAgentCallback" yaml:"beforeAgentCallback"`
	// before_model_callback block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#before_model_callback CesGuardrail#before_model_callback}
	BeforeModelCallback *CesGuardrailCodeCallbackBeforeModelCallback `field:"optional" json:"beforeModelCallback" yaml:"beforeModelCallback"`
}

