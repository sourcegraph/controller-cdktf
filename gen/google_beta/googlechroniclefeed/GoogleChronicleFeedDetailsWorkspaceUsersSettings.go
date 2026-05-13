package googlechroniclefeed


type GoogleChronicleFeedDetailsWorkspaceUsersSettings struct {
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#authentication GoogleChronicleFeed#authentication}
	Authentication *GoogleChronicleFeedDetailsWorkspaceUsersSettingsAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// Projection Type. Possible values: BASIC_PROJECTION FULL_PROJECTION.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#projection_type GoogleChronicleFeed#projection_type}
	ProjectionType *string `field:"optional" json:"projectionType" yaml:"projectionType"`
	// Customer ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#workspace_customer_id GoogleChronicleFeed#workspace_customer_id}
	WorkspaceCustomerId *string `field:"optional" json:"workspaceCustomerId" yaml:"workspaceCustomerId"`
}

