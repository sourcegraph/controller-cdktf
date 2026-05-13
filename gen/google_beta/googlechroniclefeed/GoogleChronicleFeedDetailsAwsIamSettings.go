package googlechroniclefeed


type GoogleChronicleFeedDetailsAwsIamSettings struct {
	// Supported AWS IAM api type. Possible values: USERS ROLES GROUPS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#api_type GoogleChronicleFeed#api_type}
	ApiType *string `field:"optional" json:"apiType" yaml:"apiType"`
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#authentication GoogleChronicleFeed#authentication}
	Authentication *GoogleChronicleFeedDetailsAwsIamSettingsAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
}

