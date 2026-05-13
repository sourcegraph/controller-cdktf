package googlecesguardrail


type GoogleCesGuardrailAction struct {
	// generative_answer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_guardrail#generative_answer GoogleCesGuardrail#generative_answer}
	GenerativeAnswer *GoogleCesGuardrailActionGenerativeAnswer `field:"optional" json:"generativeAnswer" yaml:"generativeAnswer"`
	// respond_immediately block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_guardrail#respond_immediately GoogleCesGuardrail#respond_immediately}
	RespondImmediately *GoogleCesGuardrailActionRespondImmediately `field:"optional" json:"respondImmediately" yaml:"respondImmediately"`
	// transfer_agent block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_guardrail#transfer_agent GoogleCesGuardrail#transfer_agent}
	TransferAgent *GoogleCesGuardrailActionTransferAgent `field:"optional" json:"transferAgent" yaml:"transferAgent"`
}

