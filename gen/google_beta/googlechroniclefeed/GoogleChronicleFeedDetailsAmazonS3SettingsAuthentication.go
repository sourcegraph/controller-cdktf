package googlechroniclefeed


type GoogleChronicleFeedDetailsAmazonS3SettingsAuthentication struct {
	// Possible values: US_EAST_1 US_EAST_2 US_WEST_1 US_WEST_2 US_GOV_CLOUD US_GOV_EAST_1 EU_WEST_1 EU_WEST_2 EU_WEST_3 EU_CENTRAL_1 EU_NORTH_1 EU_SOUTH_1 AP_SOUTH_1 AP_SOUTHEAST_1 AP_SOUTHEAST_2 AP_SOUTHEAST_3 AP_NORTHEAST_1 AP_NORTHEAST_2 AP_NORTHEAST_3 AP_EAST_1 SA_EAST_1 CN_NORTH_1 CN_NORTHWEST_1 CA_CENTRAL_1 AF_SOUTH_1 ME_SOUTH_1 AP_SOUTH_2 AP_SOUTHEAST_4 CA_WEST_1 EU_SOUTH_2 EU_CENTRAL_2 IL_CENTRAL_1 ME_CENTRAL_1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#region GoogleChronicleFeed#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// Access key ID. Used when using access key auth.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#access_key_id GoogleChronicleFeed#access_key_id}
	AccessKeyId *string `field:"optional" json:"accessKeyId" yaml:"accessKeyId"`
	// Client ID. Used when using OAuth auth.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#client_id GoogleChronicleFeed#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Client secret. Used when using OAuth auth.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#client_secret GoogleChronicleFeed#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Refresh URI. Used when using OAuth auth.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#refresh_uri GoogleChronicleFeed#refresh_uri}
	RefreshUri *string `field:"optional" json:"refreshUri" yaml:"refreshUri"`
	// Secret access key. Used when using access key auth.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#secret_access_key GoogleChronicleFeed#secret_access_key}
	SecretAccessKey *string `field:"optional" json:"secretAccessKey" yaml:"secretAccessKey"`
}

