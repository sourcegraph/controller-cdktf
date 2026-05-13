package googlemodelarmortemplate


type GoogleModelArmorTemplateFilterConfigSdpSettingsBasicConfig struct {
	// Tells whether the Sensitive Data Protection basic config is enabled or disabled. Possible values: ENABLED DISABLED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_model_armor_template#filter_enforcement GoogleModelArmorTemplate#filter_enforcement}
	FilterEnforcement *string `field:"optional" json:"filterEnforcement" yaml:"filterEnforcement"`
}

