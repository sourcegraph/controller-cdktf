package googlefirebaseailogicconfig


type GoogleFirebaseAiLogicConfigGenerativeLanguageConfig struct {
	// The value of the API key.
	//
	// The API key must have
	// 'generativelanguage.googleapis.com' in its "API restrictions" allowlist.
	// Note that this API is sometimes called the *Generative Language API* in
	// the Google Cloud console.
	//
	// Do **not** add this Gemini API key into your app's codebase
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_ai_logic_config#api_key GoogleFirebaseAiLogicConfig#api_key}
	ApiKey *string `field:"optional" json:"apiKey" yaml:"apiKey"`
	// The value of the API key.
	//
	// The API key must have
	// 'generativelanguage.googleapis.com' in its "API restrictions" allowlist.
	// Note that this API is sometimes called the *Generative Language API* in
	// the Google Cloud console.
	//
	// Do **not** add this Gemini API key into your app's codebase
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_ai_logic_config#api_key_wo GoogleFirebaseAiLogicConfig#api_key_wo}
	ApiKeyWo *string `field:"optional" json:"apiKeyWo" yaml:"apiKeyWo"`
	// Triggers update of 'api_key_wo' write-only.
	//
	// Increment this value when an update to 'api_key_wo' is needed. For more info see [updating write-only arguments](/docs/providers/google/guides/using_write_only_arguments.html#updating-write-only-arguments)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_ai_logic_config#api_key_wo_version GoogleFirebaseAiLogicConfig#api_key_wo_version}
	ApiKeyWoVersion *string `field:"optional" json:"apiKeyWoVersion" yaml:"apiKeyWoVersion"`
}

