package googleiapsettings


type GoogleIapSettingsAccessSettingsCorsSettings struct {
	// Configuration to allow HTTP OPTIONS calls to skip authorization.
	//
	// If undefined, IAP will not apply any special logic to OPTIONS requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_iap_settings#allow_http_options GoogleIapSettings#allow_http_options}
	AllowHttpOptions interface{} `field:"optional" json:"allowHttpOptions" yaml:"allowHttpOptions"`
}

