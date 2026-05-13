package googlemodelarmorfloorsetting


type GoogleModelArmorFloorsettingFilterConfigPiAndJailbreakFilterSettings struct {
	// Possible values: LOW_AND_ABOVE MEDIUM_AND_ABOVE HIGH.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_model_armor_floorsetting#confidence_level GoogleModelArmorFloorsetting#confidence_level}
	ConfidenceLevel *string `field:"optional" json:"confidenceLevel" yaml:"confidenceLevel"`
	// Tells whether Prompt injection and Jailbreak filter is enabled or disabled. Possible values: ENABLED DISABLED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_model_armor_floorsetting#filter_enforcement GoogleModelArmorFloorsetting#filter_enforcement}
	FilterEnforcement *string `field:"optional" json:"filterEnforcement" yaml:"filterEnforcement"`
}

