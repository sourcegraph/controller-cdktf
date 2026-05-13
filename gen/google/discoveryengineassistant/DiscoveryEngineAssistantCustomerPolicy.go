package discoveryengineassistant


type DiscoveryEngineAssistantCustomerPolicy struct {
	// banned_phrases block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_assistant#banned_phrases DiscoveryEngineAssistant#banned_phrases}
	BannedPhrases interface{} `field:"optional" json:"bannedPhrases" yaml:"bannedPhrases"`
	// model_armor_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_assistant#model_armor_config DiscoveryEngineAssistant#model_armor_config}
	ModelArmorConfig *DiscoveryEngineAssistantCustomerPolicyModelArmorConfig `field:"optional" json:"modelArmorConfig" yaml:"modelArmorConfig"`
}

