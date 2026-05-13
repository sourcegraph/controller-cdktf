package googlechroniclefeed


type GoogleChronicleFeedDetailsOktaUserContextSettings struct {
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#authentication GoogleChronicleFeed#authentication}
	Authentication *GoogleChronicleFeedDetailsOktaUserContextSettingsAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// API Hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#hostname GoogleChronicleFeed#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Manager id reference field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#manager_id_reference_field GoogleChronicleFeed#manager_id_reference_field}
	ManagerIdReferenceField *string `field:"optional" json:"managerIdReferenceField" yaml:"managerIdReferenceField"`
}

