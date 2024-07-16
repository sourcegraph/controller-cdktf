package googlesecurityscannerscanconfig


type GoogleSecurityScannerScanConfigAuthenticationGoogleAccount struct {
	// The password of the Google account. The credential is stored encrypted in GCP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_security_scanner_scan_config#password GoogleSecurityScannerScanConfig#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// The user name of the Google account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_security_scanner_scan_config#username GoogleSecurityScannerScanConfig#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

