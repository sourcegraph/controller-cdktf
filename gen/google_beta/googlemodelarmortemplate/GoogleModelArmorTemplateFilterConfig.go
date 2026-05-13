package googlemodelarmortemplate


type GoogleModelArmorTemplateFilterConfig struct {
	// malicious_uri_filter_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_model_armor_template#malicious_uri_filter_settings GoogleModelArmorTemplate#malicious_uri_filter_settings}
	MaliciousUriFilterSettings *GoogleModelArmorTemplateFilterConfigMaliciousUriFilterSettings `field:"optional" json:"maliciousUriFilterSettings" yaml:"maliciousUriFilterSettings"`
	// pi_and_jailbreak_filter_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_model_armor_template#pi_and_jailbreak_filter_settings GoogleModelArmorTemplate#pi_and_jailbreak_filter_settings}
	PiAndJailbreakFilterSettings *GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettings `field:"optional" json:"piAndJailbreakFilterSettings" yaml:"piAndJailbreakFilterSettings"`
	// rai_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_model_armor_template#rai_settings GoogleModelArmorTemplate#rai_settings}
	RaiSettings *GoogleModelArmorTemplateFilterConfigRaiSettings `field:"optional" json:"raiSettings" yaml:"raiSettings"`
	// sdp_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_model_armor_template#sdp_settings GoogleModelArmorTemplate#sdp_settings}
	SdpSettings *GoogleModelArmorTemplateFilterConfigSdpSettings `field:"optional" json:"sdpSettings" yaml:"sdpSettings"`
}

