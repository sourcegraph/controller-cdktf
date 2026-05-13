package googlecesapp


type GoogleCesAppAudioProcessingConfigSynthesizeSpeechConfigs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#language_code GoogleCesApp#language_code}.
	LanguageCode *string `field:"required" json:"languageCode" yaml:"languageCode"`
	// The speaking rate/speed in the range [0.25, 2.0]. 1.0 is the normal native speed supported by the specific voice. 2.0 is twice as fast, and 0.5 is half as fast. Values outside of the range [0.25, 2.0] will return an error.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#speaking_rate GoogleCesApp#speaking_rate}
	SpeakingRate *float64 `field:"optional" json:"speakingRate" yaml:"speakingRate"`
	// The name of the voice.
	//
	// If not set, the service will choose a
	// voice based on the other parameters such as language_code.
	// For the list of available voices, please refer to Supported voices and
	// languages from Cloud Text-to-Speech.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#voice GoogleCesApp#voice}
	Voice *string `field:"optional" json:"voice" yaml:"voice"`
}

