package cesguardrail


type CesGuardrailModelSafetySafetySettings struct {
	// The harm category. Possible values: HARM_CATEGORY_HATE_SPEECH HARM_CATEGORY_DANGEROUS_CONTENT HARM_CATEGORY_HARASSMENT HARM_CATEGORY_SEXUALLY_EXPLICIT Possible values: ["HARM_CATEGORY_HATE_SPEECH", "HARM_CATEGORY_DANGEROUS_CONTENT", "HARM_CATEGORY_HARASSMENT", "HARM_CATEGORY_SEXUALLY_EXPLICIT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#category CesGuardrail#category}
	Category *string `field:"required" json:"category" yaml:"category"`
	// The harm block threshold. Possible values: BLOCK_LOW_AND_ABOVE BLOCK_MEDIUM_AND_ABOVE BLOCK_ONLY_HIGH BLOCK_NONE OFF Possible values: ["BLOCK_LOW_AND_ABOVE", "BLOCK_MEDIUM_AND_ABOVE", "BLOCK_ONLY_HIGH", "BLOCK_NONE", "OFF"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#threshold CesGuardrail#threshold}
	Threshold *string `field:"required" json:"threshold" yaml:"threshold"`
}

