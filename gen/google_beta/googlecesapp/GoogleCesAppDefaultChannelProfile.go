package googlecesapp


type GoogleCesAppDefaultChannelProfile struct {
	// The type of the channel profile. Possible values: UNKNOWN WEB_UI API TWILIO GOOGLE_TELEPHONY_PLATFORM CONTACT_CENTER_AS_A_SERVICE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#channel_type GoogleCesApp#channel_type}
	ChannelType *string `field:"optional" json:"channelType" yaml:"channelType"`
	// Whether to disable user barge-in in the conversation.
	//
	// - true: User interruptions are disabled while the agent is speaking.
	// - false: The agent retains automatic control over when the user can interrupt.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#disable_barge_in_control GoogleCesApp#disable_barge_in_control}
	DisableBargeInControl interface{} `field:"optional" json:"disableBargeInControl" yaml:"disableBargeInControl"`
	// Whether to disable DTMF (dual-tone multi-frequency).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#disable_dtmf GoogleCesApp#disable_dtmf}
	DisableDtmf interface{} `field:"optional" json:"disableDtmf" yaml:"disableDtmf"`
	// persona_property block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#persona_property GoogleCesApp#persona_property}
	PersonaProperty *GoogleCesAppDefaultChannelProfilePersonaProperty `field:"optional" json:"personaProperty" yaml:"personaProperty"`
	// The unique identifier of the channel profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#profile_id GoogleCesApp#profile_id}
	ProfileId *string `field:"optional" json:"profileId" yaml:"profileId"`
	// web_widget_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#web_widget_config GoogleCesApp#web_widget_config}
	WebWidgetConfig *GoogleCesAppDefaultChannelProfileWebWidgetConfig `field:"optional" json:"webWidgetConfig" yaml:"webWidgetConfig"`
}

