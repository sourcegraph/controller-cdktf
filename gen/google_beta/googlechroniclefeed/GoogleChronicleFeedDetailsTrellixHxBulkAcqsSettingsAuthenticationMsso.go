package googlechroniclefeed


type GoogleChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationMsso struct {
	// The login api endpoint url.
	//
	// This must be a valid URL with an http or https scheme. It has no default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#api_endpoint GoogleChronicleFeed#api_endpoint}
	ApiEndpoint *string `field:"required" json:"apiEndpoint" yaml:"apiEndpoint"`
	// Password of the account identified by username.
	//
	// There are no restrictions on the format of the password. It has no default,
	// specifically enforced min / max length or character set. The password
	// will have been provided by an MSSO administrator and it is assumed that
	// they have provided a password that is internally consistent with MSSO
	// authentication requirements / validation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#password GoogleChronicleFeed#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// Username for MSSO authentication.
	//
	// There are no restrictions on the format of the username. It has no default,
	// specifically enforced min / max length or character set. The username
	// will have been provided by an MSSO administrator and it is assumed that
	// they have provided a username that is internally consistent with MSSO
	// authentication requirements / validation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#username GoogleChronicleFeed#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

