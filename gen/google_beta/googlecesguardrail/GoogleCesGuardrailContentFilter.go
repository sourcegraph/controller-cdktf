package googlecesguardrail


type GoogleCesGuardrailContentFilter struct {
	// Match type for the content filter. Possible values: SIMPLE_STRING_MATCH WORD_BOUNDARY_STRING_MATCH REGEXP_MATCH.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_guardrail#match_type GoogleCesGuardrail#match_type}
	MatchType *string `field:"required" json:"matchType" yaml:"matchType"`
	// List of banned phrases. Applies to both user inputs and agent responses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_guardrail#banned_contents GoogleCesGuardrail#banned_contents}
	BannedContents *[]*string `field:"optional" json:"bannedContents" yaml:"bannedContents"`
	// List of banned phrases. Applies only to agent responses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_guardrail#banned_contents_in_agent_response GoogleCesGuardrail#banned_contents_in_agent_response}
	BannedContentsInAgentResponse *[]*string `field:"optional" json:"bannedContentsInAgentResponse" yaml:"bannedContentsInAgentResponse"`
	// List of banned phrases. Applies only to user inputs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_guardrail#banned_contents_in_user_input GoogleCesGuardrail#banned_contents_in_user_input}
	BannedContentsInUserInput *[]*string `field:"optional" json:"bannedContentsInUserInput" yaml:"bannedContentsInUserInput"`
	// If true, diacritics are ignored during matching.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_guardrail#disregard_diacritics GoogleCesGuardrail#disregard_diacritics}
	DisregardDiacritics interface{} `field:"optional" json:"disregardDiacritics" yaml:"disregardDiacritics"`
}

