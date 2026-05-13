package cesapp


type CesAppLanguageSettings struct {
	// The default language code of the app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#default_language_code CesApp#default_language_code}
	DefaultLanguageCode *string `field:"optional" json:"defaultLanguageCode" yaml:"defaultLanguageCode"`
	// Enables multilingual support. If true, agents in the app will use pre-built instructions to improve handling of multilingual input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#enable_multilingual_support CesApp#enable_multilingual_support}
	EnableMultilingualSupport interface{} `field:"optional" json:"enableMultilingualSupport" yaml:"enableMultilingualSupport"`
	// The action to perform when an agent receives input in an unsupported language.
	//
	// This can be a predefined action or a custom tool call.
	// Valid values are:
	// - A tool's full resource name, which triggers a specific tool execution.
	// - A predefined system action, such as "escalate" or "exit", which triggers
	// an EndSession signal with corresponding metadata
	// to terminate the conversation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#fallback_action CesApp#fallback_action}
	FallbackAction *string `field:"optional" json:"fallbackAction" yaml:"fallbackAction"`
	// List of languages codes supported by the app, in addition to the 'default_language_code'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#supported_language_codes CesApp#supported_language_codes}
	SupportedLanguageCodes *[]*string `field:"optional" json:"supportedLanguageCodes" yaml:"supportedLanguageCodes"`
}

