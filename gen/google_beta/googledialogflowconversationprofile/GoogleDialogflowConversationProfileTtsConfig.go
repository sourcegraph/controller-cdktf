package googledialogflowconversationprofile


type GoogleDialogflowConversationProfileTtsConfig struct {
	// An identifier which selects 'audio effects' profiles that are applied on (post synthesized) text to speech.
	//
	// Effects are applied on top of each other in the order they are given.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#effects_profile_id GoogleDialogflowConversationProfile#effects_profile_id}
	EffectsProfileId *[]*string `field:"optional" json:"effectsProfileId" yaml:"effectsProfileId"`
	// Speaking pitch, in the range [-20.0, 20.0]. 20 means increase 20 semitones from the original pitch. -20 means decrease 20 semitones from the original pitch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#pitch GoogleDialogflowConversationProfile#pitch}
	Pitch *float64 `field:"optional" json:"pitch" yaml:"pitch"`
	// Speaking rate/speed, in the range [0.25, 4.0].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#speaking_rate GoogleDialogflowConversationProfile#speaking_rate}
	SpeakingRate *float64 `field:"optional" json:"speakingRate" yaml:"speakingRate"`
	// voice block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#voice GoogleDialogflowConversationProfile#voice}
	Voice *GoogleDialogflowConversationProfileTtsConfigVoice `field:"optional" json:"voice" yaml:"voice"`
	// Volume gain (in dB) of the normal native volume supported by the specific voice.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#volume_gain_db GoogleDialogflowConversationProfile#volume_gain_db}
	VolumeGainDb *float64 `field:"optional" json:"volumeGainDb" yaml:"volumeGainDb"`
}

