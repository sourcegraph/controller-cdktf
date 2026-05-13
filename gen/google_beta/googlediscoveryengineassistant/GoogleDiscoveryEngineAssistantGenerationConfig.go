package googlediscoveryengineassistant


type GoogleDiscoveryEngineAssistantGenerationConfig struct {
	// The default language to use for the generation of the assistant response.
	//
	// Use an ISO 639-1 language code such as 'en'.
	// If not specified, the language will be automatically detected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_assistant#default_language GoogleDiscoveryEngineAssistant#default_language}
	DefaultLanguage *string `field:"optional" json:"defaultLanguage" yaml:"defaultLanguage"`
	// system_instruction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_assistant#system_instruction GoogleDiscoveryEngineAssistant#system_instruction}
	SystemInstruction *GoogleDiscoveryEngineAssistantGenerationConfigSystemInstruction `field:"optional" json:"systemInstruction" yaml:"systemInstruction"`
}

