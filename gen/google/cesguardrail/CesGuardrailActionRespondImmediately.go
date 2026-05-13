package cesguardrail


type CesGuardrailActionRespondImmediately struct {
	// responses block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#responses CesGuardrail#responses}
	Responses interface{} `field:"required" json:"responses" yaml:"responses"`
}

