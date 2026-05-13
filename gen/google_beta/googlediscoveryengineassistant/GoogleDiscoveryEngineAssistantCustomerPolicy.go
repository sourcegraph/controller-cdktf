package googlediscoveryengineassistant


type GoogleDiscoveryEngineAssistantCustomerPolicy struct {
	// banned_phrases block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_assistant#banned_phrases GoogleDiscoveryEngineAssistant#banned_phrases}
	BannedPhrases interface{} `field:"optional" json:"bannedPhrases" yaml:"bannedPhrases"`
	// model_armor_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_assistant#model_armor_config GoogleDiscoveryEngineAssistant#model_armor_config}
	ModelArmorConfig *GoogleDiscoveryEngineAssistantCustomerPolicyModelArmorConfig `field:"optional" json:"modelArmorConfig" yaml:"modelArmorConfig"`
}

