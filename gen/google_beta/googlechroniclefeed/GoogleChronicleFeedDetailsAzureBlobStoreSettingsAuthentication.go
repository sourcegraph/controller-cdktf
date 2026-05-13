package googlechroniclefeed


type GoogleChronicleFeedDetailsAzureBlobStoreSettingsAuthentication struct {
	// SAS Token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#sas_token GoogleChronicleFeed#sas_token}
	SasToken *string `field:"optional" json:"sasToken" yaml:"sasToken"`
	// Shared Key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#shared_key GoogleChronicleFeed#shared_key}
	SharedKey *string `field:"optional" json:"sharedKey" yaml:"sharedKey"`
}

