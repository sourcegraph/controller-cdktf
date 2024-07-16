package googleapikeyskey


type GoogleApikeysKeyRestrictions struct {
	// android_key_restrictions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apikeys_key#android_key_restrictions GoogleApikeysKey#android_key_restrictions}
	AndroidKeyRestrictions *GoogleApikeysKeyRestrictionsAndroidKeyRestrictions `field:"optional" json:"androidKeyRestrictions" yaml:"androidKeyRestrictions"`
	// api_targets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apikeys_key#api_targets GoogleApikeysKey#api_targets}
	ApiTargets interface{} `field:"optional" json:"apiTargets" yaml:"apiTargets"`
	// browser_key_restrictions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apikeys_key#browser_key_restrictions GoogleApikeysKey#browser_key_restrictions}
	BrowserKeyRestrictions *GoogleApikeysKeyRestrictionsBrowserKeyRestrictions `field:"optional" json:"browserKeyRestrictions" yaml:"browserKeyRestrictions"`
	// ios_key_restrictions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apikeys_key#ios_key_restrictions GoogleApikeysKey#ios_key_restrictions}
	IosKeyRestrictions *GoogleApikeysKeyRestrictionsIosKeyRestrictions `field:"optional" json:"iosKeyRestrictions" yaml:"iosKeyRestrictions"`
	// server_key_restrictions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apikeys_key#server_key_restrictions GoogleApikeysKey#server_key_restrictions}
	ServerKeyRestrictions *GoogleApikeysKeyRestrictionsServerKeyRestrictions `field:"optional" json:"serverKeyRestrictions" yaml:"serverKeyRestrictions"`
}

