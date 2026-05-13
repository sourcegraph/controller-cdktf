package googlechroniclefeed


type GoogleChronicleFeedDetailsTrellixHxHostsSettingsAuthentication struct {
	// msso block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#msso GoogleChronicleFeed#msso}
	Msso *GoogleChronicleFeedDetailsTrellixHxHostsSettingsAuthenticationMsso `field:"optional" json:"msso" yaml:"msso"`
	// trellix_iam block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#trellix_iam GoogleChronicleFeed#trellix_iam}
	TrellixIam *GoogleChronicleFeedDetailsTrellixHxHostsSettingsAuthenticationTrellixIam `field:"optional" json:"trellixIam" yaml:"trellixIam"`
}

