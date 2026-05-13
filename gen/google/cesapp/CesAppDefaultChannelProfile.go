package cesapp


type CesAppDefaultChannelProfile struct {
	// The type of the channel profile. Possible values: UNKNOWN WEB_UI API TWILIO GOOGLE_TELEPHONY_PLATFORM CONTACT_CENTER_AS_A_SERVICE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#channel_type CesApp#channel_type}
	ChannelType *string `field:"optional" json:"channelType" yaml:"channelType"`
	// Whether to disable user barge-in in the conversation.
	//
	// - true: User interruptions are disabled while the agent is speaking.
	// - false: The agent retains automatic control over when the user can interrupt.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#disable_barge_in_control CesApp#disable_barge_in_control}
	DisableBargeInControl interface{} `field:"optional" json:"disableBargeInControl" yaml:"disableBargeInControl"`
	// Whether to disable DTMF (dual-tone multi-frequency).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#disable_dtmf CesApp#disable_dtmf}
	DisableDtmf interface{} `field:"optional" json:"disableDtmf" yaml:"disableDtmf"`
	// persona_property block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#persona_property CesApp#persona_property}
	PersonaProperty *CesAppDefaultChannelProfilePersonaProperty `field:"optional" json:"personaProperty" yaml:"personaProperty"`
	// The unique identifier of the channel profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#profile_id CesApp#profile_id}
	ProfileId *string `field:"optional" json:"profileId" yaml:"profileId"`
	// web_widget_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#web_widget_config CesApp#web_widget_config}
	WebWidgetConfig *CesAppDefaultChannelProfileWebWidgetConfig `field:"optional" json:"webWidgetConfig" yaml:"webWidgetConfig"`
}

