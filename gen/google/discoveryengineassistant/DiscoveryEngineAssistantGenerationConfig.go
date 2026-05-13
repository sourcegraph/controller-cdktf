package discoveryengineassistant


type DiscoveryEngineAssistantGenerationConfig struct {
	// The default language to use for the generation of the assistant response.
	//
	// Use an ISO 639-1 language code such as 'en'.
	// If not specified, the language will be automatically detected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_assistant#default_language DiscoveryEngineAssistant#default_language}
	DefaultLanguage *string `field:"optional" json:"defaultLanguage" yaml:"defaultLanguage"`
	// system_instruction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_assistant#system_instruction DiscoveryEngineAssistant#system_instruction}
	SystemInstruction *DiscoveryEngineAssistantGenerationConfigSystemInstruction `field:"optional" json:"systemInstruction" yaml:"systemInstruction"`
}

