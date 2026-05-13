package googlechroniclefeed


type GoogleChronicleFeedDetailsPanPrismaCloudSettingsAuthentication struct {
	// Password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#password GoogleChronicleFeed#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Username.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#user GoogleChronicleFeed#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}

