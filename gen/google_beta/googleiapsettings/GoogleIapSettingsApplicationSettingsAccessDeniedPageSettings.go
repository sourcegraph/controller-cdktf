package googleiapsettings


type GoogleIapSettingsApplicationSettingsAccessDeniedPageSettings struct {
	// The URI to be redirected to when access is denied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_iap_settings#access_denied_page_uri GoogleIapSettings#access_denied_page_uri}
	AccessDeniedPageUri *string `field:"optional" json:"accessDeniedPageUri" yaml:"accessDeniedPageUri"`
	// Whether to generate a troubleshooting URL on access denied events to this application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_iap_settings#generate_troubleshooting_uri GoogleIapSettings#generate_troubleshooting_uri}
	GenerateTroubleshootingUri interface{} `field:"optional" json:"generateTroubleshootingUri" yaml:"generateTroubleshootingUri"`
	// Whether to generate remediation token on access denied events to this application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_iap_settings#remediation_token_generation_enabled GoogleIapSettings#remediation_token_generation_enabled}
	RemediationTokenGenerationEnabled interface{} `field:"optional" json:"remediationTokenGenerationEnabled" yaml:"remediationTokenGenerationEnabled"`
}

