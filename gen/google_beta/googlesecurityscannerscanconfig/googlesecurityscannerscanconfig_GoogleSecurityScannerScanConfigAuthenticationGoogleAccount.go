package googlesecurityscannerscanconfig


type GoogleSecurityScannerScanConfigAuthenticationGoogleAccount struct {
	// The password of the Google account. The credential is stored encrypted in GCP.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_security_scanner_scan_config#password GoogleSecurityScannerScanConfig#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// The user name of the Google account.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_security_scanner_scan_config#username GoogleSecurityScannerScanConfig#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

