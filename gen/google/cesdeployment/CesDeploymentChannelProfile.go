package cesdeployment


type CesDeploymentChannelProfile struct {
	// The type of the channel profile. Possible values: UNKNOWN WEB_UI API TWILIO GOOGLE_TELEPHONY_PLATFORM CONTACT_CENTER_AS_A_SERVICE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_deployment#channel_type CesDeployment#channel_type}
	ChannelType *string `field:"optional" json:"channelType" yaml:"channelType"`
	// Whether to disable user barge-in control in the conversation.
	//
	// - **true**: User interruptions are disabled while the agent is speaking.
	// - **false**: The agent retains automatic control over when the user can
	// interrupt.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_deployment#disable_barge_in_control CesDeployment#disable_barge_in_control}
	DisableBargeInControl interface{} `field:"optional" json:"disableBargeInControl" yaml:"disableBargeInControl"`
	// Whether to disable DTMF (dual-tone multi-frequency).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_deployment#disable_dtmf CesDeployment#disable_dtmf}
	DisableDtmf interface{} `field:"optional" json:"disableDtmf" yaml:"disableDtmf"`
	// persona_property block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_deployment#persona_property CesDeployment#persona_property}
	PersonaProperty *CesDeploymentChannelProfilePersonaProperty `field:"optional" json:"personaProperty" yaml:"personaProperty"`
	// The unique identifier of the channel profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_deployment#profile_id CesDeployment#profile_id}
	ProfileId *string `field:"optional" json:"profileId" yaml:"profileId"`
	// web_widget_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_deployment#web_widget_config CesDeployment#web_widget_config}
	WebWidgetConfig *CesDeploymentChannelProfileWebWidgetConfig `field:"optional" json:"webWidgetConfig" yaml:"webWidgetConfig"`
}

