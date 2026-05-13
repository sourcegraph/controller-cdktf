package googlechroniclefeed


type GoogleChronicleFeedDetailsCloudPassageSettingsAuthentication struct {
	// Secret of the account identified by user_name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#secret GoogleChronicleFeed#secret}
	Secret *string `field:"optional" json:"secret" yaml:"secret"`
	// Username of an identity used for authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#user GoogleChronicleFeed#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}

