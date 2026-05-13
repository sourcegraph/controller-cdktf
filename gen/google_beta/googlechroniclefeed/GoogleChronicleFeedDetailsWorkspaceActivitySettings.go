package googlechroniclefeed


type GoogleChronicleFeedDetailsWorkspaceActivitySettings struct {
	// Applications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#applications GoogleChronicleFeed#applications}
	Applications *[]*string `field:"optional" json:"applications" yaml:"applications"`
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#authentication GoogleChronicleFeed#authentication}
	Authentication *GoogleChronicleFeedDetailsWorkspaceActivitySettingsAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// Customer ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#workspace_customer_id GoogleChronicleFeed#workspace_customer_id}
	WorkspaceCustomerId *string `field:"optional" json:"workspaceCustomerId" yaml:"workspaceCustomerId"`
}

