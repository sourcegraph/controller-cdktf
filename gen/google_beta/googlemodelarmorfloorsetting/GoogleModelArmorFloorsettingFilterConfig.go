package googlemodelarmorfloorsetting


type GoogleModelArmorFloorsettingFilterConfig struct {
	// malicious_uri_filter_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_model_armor_floorsetting#malicious_uri_filter_settings GoogleModelArmorFloorsetting#malicious_uri_filter_settings}
	MaliciousUriFilterSettings *GoogleModelArmorFloorsettingFilterConfigMaliciousUriFilterSettings `field:"optional" json:"maliciousUriFilterSettings" yaml:"maliciousUriFilterSettings"`
	// pi_and_jailbreak_filter_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_model_armor_floorsetting#pi_and_jailbreak_filter_settings GoogleModelArmorFloorsetting#pi_and_jailbreak_filter_settings}
	PiAndJailbreakFilterSettings *GoogleModelArmorFloorsettingFilterConfigPiAndJailbreakFilterSettings `field:"optional" json:"piAndJailbreakFilterSettings" yaml:"piAndJailbreakFilterSettings"`
	// rai_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_model_armor_floorsetting#rai_settings GoogleModelArmorFloorsetting#rai_settings}
	RaiSettings *GoogleModelArmorFloorsettingFilterConfigRaiSettings `field:"optional" json:"raiSettings" yaml:"raiSettings"`
	// sdp_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_model_armor_floorsetting#sdp_settings GoogleModelArmorFloorsetting#sdp_settings}
	SdpSettings *GoogleModelArmorFloorsettingFilterConfigSdpSettings `field:"optional" json:"sdpSettings" yaml:"sdpSettings"`
}

