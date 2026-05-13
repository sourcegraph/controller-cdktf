package googlemodelarmortemplate


type GoogleModelArmorTemplateFilterConfigSdpSettings struct {
	// advanced_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_model_armor_template#advanced_config GoogleModelArmorTemplate#advanced_config}
	AdvancedConfig *GoogleModelArmorTemplateFilterConfigSdpSettingsAdvancedConfig `field:"optional" json:"advancedConfig" yaml:"advancedConfig"`
	// basic_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_model_armor_template#basic_config GoogleModelArmorTemplate#basic_config}
	BasicConfig *GoogleModelArmorTemplateFilterConfigSdpSettingsBasicConfig `field:"optional" json:"basicConfig" yaml:"basicConfig"`
}

