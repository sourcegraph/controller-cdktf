package googlecesdeployment


type GoogleCesDeploymentChannelProfileWebWidgetConfig struct {
	// The modality of the web widget. Possible values: UNKNOWN_MODALITY CHAT_AND_VOICE VOICE_ONLY CHAT_ONLY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_deployment#modality GoogleCesDeployment#modality}
	Modality *string `field:"optional" json:"modality" yaml:"modality"`
	// The theme of the web widget. Possible values: UNKNOWN_THEME LIGHT DARK.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_deployment#theme GoogleCesDeployment#theme}
	Theme *string `field:"optional" json:"theme" yaml:"theme"`
	// The title of the web widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_deployment#web_widget_title GoogleCesDeployment#web_widget_title}
	WebWidgetTitle *string `field:"optional" json:"webWidgetTitle" yaml:"webWidgetTitle"`
}

