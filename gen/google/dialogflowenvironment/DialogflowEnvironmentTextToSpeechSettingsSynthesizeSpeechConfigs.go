package dialogflowenvironment


type DialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_environment#language DialogflowEnvironment#language}.
	Language *string `field:"required" json:"language" yaml:"language"`
	// An identifier which selects 'audio effects' profiles that are applied on (post synthesized) text to speech.
	//
	// Effects are applied on top of each other in the order they are given.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_environment#effects_profile_id DialogflowEnvironment#effects_profile_id}
	EffectsProfileId *[]*string `field:"optional" json:"effectsProfileId" yaml:"effectsProfileId"`
	// Speaking pitch, in the range [-20.0, 20.0]. 20 means increase 20 semitones from the original pitch. -20 means decrease 20 semitones from the original pitch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_environment#pitch DialogflowEnvironment#pitch}
	Pitch *float64 `field:"optional" json:"pitch" yaml:"pitch"`
	// Speaking rate/speed, in the range [0.25, 4.0].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_environment#speaking_rate DialogflowEnvironment#speaking_rate}
	SpeakingRate *float64 `field:"optional" json:"speakingRate" yaml:"speakingRate"`
	// voice block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_environment#voice DialogflowEnvironment#voice}
	Voice *DialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsVoice `field:"optional" json:"voice" yaml:"voice"`
	// Volume gain (in dB) of the normal native volume supported by the specific voice.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_environment#volume_gain_db DialogflowEnvironment#volume_gain_db}
	VolumeGainDb *float64 `field:"optional" json:"volumeGainDb" yaml:"volumeGainDb"`
}

