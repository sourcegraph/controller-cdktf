package googlediscoveryengineassistant


type GoogleDiscoveryEngineAssistantCustomerPolicyBannedPhrases struct {
	// The raw string content to be banned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_assistant#phrase GoogleDiscoveryEngineAssistant#phrase}
	Phrase *string `field:"required" json:"phrase" yaml:"phrase"`
	// If true, diacritical marks (e.g., accents, umlauts) are ignored when matching banned phrases. For example, "cafe" would match "café".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_assistant#ignore_diacritics GoogleDiscoveryEngineAssistant#ignore_diacritics}
	IgnoreDiacritics interface{} `field:"optional" json:"ignoreDiacritics" yaml:"ignoreDiacritics"`
	// Match type for the banned phrase. The supported values: 'SIMPLE_STRING_MATCH', 'WORD_BOUNDARY_STRING_MATCH'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_assistant#match_type GoogleDiscoveryEngineAssistant#match_type}
	MatchType *string `field:"optional" json:"matchType" yaml:"matchType"`
}

