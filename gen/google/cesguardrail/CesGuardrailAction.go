package cesguardrail


type CesGuardrailAction struct {
	// generative_answer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#generative_answer CesGuardrail#generative_answer}
	GenerativeAnswer *CesGuardrailActionGenerativeAnswer `field:"optional" json:"generativeAnswer" yaml:"generativeAnswer"`
	// respond_immediately block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#respond_immediately CesGuardrail#respond_immediately}
	RespondImmediately *CesGuardrailActionRespondImmediately `field:"optional" json:"respondImmediately" yaml:"respondImmediately"`
	// transfer_agent block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#transfer_agent CesGuardrail#transfer_agent}
	TransferAgent *CesGuardrailActionTransferAgent `field:"optional" json:"transferAgent" yaml:"transferAgent"`
}

