package cesapp


type CesAppAudioProcessingConfig struct {
	// ambient_sound_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#ambient_sound_config CesApp#ambient_sound_config}
	AmbientSoundConfig *CesAppAudioProcessingConfigAmbientSoundConfig `field:"optional" json:"ambientSoundConfig" yaml:"ambientSoundConfig"`
	// barge_in_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#barge_in_config CesApp#barge_in_config}
	BargeInConfig *CesAppAudioProcessingConfigBargeInConfig `field:"optional" json:"bargeInConfig" yaml:"bargeInConfig"`
	// The duration of user inactivity (no speech or interaction) before the agent prompts the user for reengagement.
	//
	// If not set, the agent will not prompt
	// the user for reengagement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#inactivity_timeout CesApp#inactivity_timeout}
	InactivityTimeout *string `field:"optional" json:"inactivityTimeout" yaml:"inactivityTimeout"`
	// synthesize_speech_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#synthesize_speech_configs CesApp#synthesize_speech_configs}
	SynthesizeSpeechConfigs interface{} `field:"optional" json:"synthesizeSpeechConfigs" yaml:"synthesizeSpeechConfigs"`
}

