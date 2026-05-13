package cesguardrail


type CesGuardrailContentFilter struct {
	// Match type for the content filter. Possible values: SIMPLE_STRING_MATCH WORD_BOUNDARY_STRING_MATCH REGEXP_MATCH.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#match_type CesGuardrail#match_type}
	MatchType *string `field:"required" json:"matchType" yaml:"matchType"`
	// List of banned phrases. Applies to both user inputs and agent responses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#banned_contents CesGuardrail#banned_contents}
	BannedContents *[]*string `field:"optional" json:"bannedContents" yaml:"bannedContents"`
	// List of banned phrases. Applies only to agent responses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#banned_contents_in_agent_response CesGuardrail#banned_contents_in_agent_response}
	BannedContentsInAgentResponse *[]*string `field:"optional" json:"bannedContentsInAgentResponse" yaml:"bannedContentsInAgentResponse"`
	// List of banned phrases. Applies only to user inputs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#banned_contents_in_user_input CesGuardrail#banned_contents_in_user_input}
	BannedContentsInUserInput *[]*string `field:"optional" json:"bannedContentsInUserInput" yaml:"bannedContentsInUserInput"`
	// If true, diacritics are ignored during matching.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_guardrail#disregard_diacritics CesGuardrail#disregard_diacritics}
	DisregardDiacritics interface{} `field:"optional" json:"disregardDiacritics" yaml:"disregardDiacritics"`
}

