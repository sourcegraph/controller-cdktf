package googlecesapp


type GoogleCesAppAudioProcessingConfig struct {
	// ambient_sound_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#ambient_sound_config GoogleCesApp#ambient_sound_config}
	AmbientSoundConfig *GoogleCesAppAudioProcessingConfigAmbientSoundConfig `field:"optional" json:"ambientSoundConfig" yaml:"ambientSoundConfig"`
	// barge_in_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#barge_in_config GoogleCesApp#barge_in_config}
	BargeInConfig *GoogleCesAppAudioProcessingConfigBargeInConfig `field:"optional" json:"bargeInConfig" yaml:"bargeInConfig"`
	// The duration of user inactivity (no speech or interaction) before the agent prompts the user for reengagement.
	//
	// If not set, the agent will not prompt
	// the user for reengagement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#inactivity_timeout GoogleCesApp#inactivity_timeout}
	InactivityTimeout *string `field:"optional" json:"inactivityTimeout" yaml:"inactivityTimeout"`
	// synthesize_speech_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#synthesize_speech_configs GoogleCesApp#synthesize_speech_configs}
	SynthesizeSpeechConfigs interface{} `field:"optional" json:"synthesizeSpeechConfigs" yaml:"synthesizeSpeechConfigs"`
}

