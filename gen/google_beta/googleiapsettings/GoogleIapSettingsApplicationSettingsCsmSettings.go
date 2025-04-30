package googleiapsettings


type GoogleIapSettingsApplicationSettingsCsmSettings struct {
	// Audience claim set in the generated RCToken. This value is not validated by IAP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_iap_settings#rctoken_aud GoogleIapSettings#rctoken_aud}
	RctokenAud *string `field:"optional" json:"rctokenAud" yaml:"rctokenAud"`
}

