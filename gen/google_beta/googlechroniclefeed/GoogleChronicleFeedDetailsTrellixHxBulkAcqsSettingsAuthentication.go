package googlechroniclefeed


type GoogleChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthentication struct {
	// msso block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#msso GoogleChronicleFeed#msso}
	Msso *GoogleChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationMsso `field:"optional" json:"msso" yaml:"msso"`
	// trellix_iam block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#trellix_iam GoogleChronicleFeed#trellix_iam}
	TrellixIam *GoogleChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationTrellixIam `field:"optional" json:"trellixIam" yaml:"trellixIam"`
}

