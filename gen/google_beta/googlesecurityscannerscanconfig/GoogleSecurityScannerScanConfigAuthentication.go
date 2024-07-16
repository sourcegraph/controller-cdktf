package googlesecurityscannerscanconfig


type GoogleSecurityScannerScanConfigAuthentication struct {
	// custom_account block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_security_scanner_scan_config#custom_account GoogleSecurityScannerScanConfig#custom_account}
	CustomAccount *GoogleSecurityScannerScanConfigAuthenticationCustomAccount `field:"optional" json:"customAccount" yaml:"customAccount"`
	// google_account block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_security_scanner_scan_config#google_account GoogleSecurityScannerScanConfig#google_account}
	GoogleAccount *GoogleSecurityScannerScanConfigAuthenticationGoogleAccount `field:"optional" json:"googleAccount" yaml:"googleAccount"`
}

