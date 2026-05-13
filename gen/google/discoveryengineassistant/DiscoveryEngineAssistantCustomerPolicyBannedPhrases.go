package discoveryengineassistant


type DiscoveryEngineAssistantCustomerPolicyBannedPhrases struct {
	// The raw string content to be banned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_assistant#phrase DiscoveryEngineAssistant#phrase}
	Phrase *string `field:"required" json:"phrase" yaml:"phrase"`
	// If true, diacritical marks (e.g., accents, umlauts) are ignored when matching banned phrases. For example, "cafe" would match "café".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_assistant#ignore_diacritics DiscoveryEngineAssistant#ignore_diacritics}
	IgnoreDiacritics interface{} `field:"optional" json:"ignoreDiacritics" yaml:"ignoreDiacritics"`
	// Match type for the banned phrase. The supported values: 'SIMPLE_STRING_MATCH', 'WORD_BOUNDARY_STRING_MATCH'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_assistant#match_type DiscoveryEngineAssistant#match_type}
	MatchType *string `field:"optional" json:"matchType" yaml:"matchType"`
}

