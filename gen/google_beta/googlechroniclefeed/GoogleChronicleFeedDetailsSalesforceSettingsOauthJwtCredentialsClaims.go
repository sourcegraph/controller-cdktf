package googlechroniclefeed


type GoogleChronicleFeedDetailsSalesforceSettingsOauthJwtCredentialsClaims struct {
	// Audience.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#audience GoogleChronicleFeed#audience}
	Audience *string `field:"optional" json:"audience" yaml:"audience"`
	// Issuer. Usually the client_id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#issuer GoogleChronicleFeed#issuer}
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// Subject. Usually the email.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#subject GoogleChronicleFeed#subject}
	Subject *string `field:"optional" json:"subject" yaml:"subject"`
}

