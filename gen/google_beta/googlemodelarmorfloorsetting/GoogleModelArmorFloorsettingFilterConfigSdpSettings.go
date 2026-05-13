package googlemodelarmorfloorsetting


type GoogleModelArmorFloorsettingFilterConfigSdpSettings struct {
	// advanced_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_model_armor_floorsetting#advanced_config GoogleModelArmorFloorsetting#advanced_config}
	AdvancedConfig *GoogleModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfig `field:"optional" json:"advancedConfig" yaml:"advancedConfig"`
	// basic_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_model_armor_floorsetting#basic_config GoogleModelArmorFloorsetting#basic_config}
	BasicConfig *GoogleModelArmorFloorsettingFilterConfigSdpSettingsBasicConfig `field:"optional" json:"basicConfig" yaml:"basicConfig"`
}

