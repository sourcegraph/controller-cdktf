package googlechroniclefeed


type GoogleChronicleFeedDetailsAmazonSqsSettings struct {
	// Account number of the owner of the queue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#account_number GoogleChronicleFeed#account_number}
	AccountNumber *string `field:"optional" json:"accountNumber" yaml:"accountNumber"`
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#authentication GoogleChronicleFeed#authentication}
	Authentication *GoogleChronicleFeedDetailsAmazonSqsSettingsAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// Name of the queue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#queue GoogleChronicleFeed#queue}
	Queue *string `field:"optional" json:"queue" yaml:"queue"`
	// Possible values: US_EAST_1 US_EAST_2 US_WEST_1 US_WEST_2 US_GOV_CLOUD US_GOV_EAST_1 EU_WEST_1 EU_WEST_2 EU_WEST_3 EU_CENTRAL_1 EU_NORTH_1 EU_SOUTH_1 AP_SOUTH_1 AP_SOUTHEAST_1 AP_SOUTHEAST_2 AP_SOUTHEAST_3 AP_NORTHEAST_1 AP_NORTHEAST_2 AP_NORTHEAST_3 AP_EAST_1 SA_EAST_1 CN_NORTH_1 CN_NORTHWEST_1 CA_CENTRAL_1 AF_SOUTH_1 ME_SOUTH_1 AP_SOUTH_2 AP_SOUTHEAST_4 CA_WEST_1 EU_SOUTH_2 EU_CENTRAL_2 IL_CENTRAL_1 ME_CENTRAL_1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#region GoogleChronicleFeed#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Possible values: SOURCE_DELETION_NEVER SOURCE_DELETION_ON_SUCCESS SOURCE_DELETION_ON_SUCCESS_FILES_ONLY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#source_deletion_option GoogleChronicleFeed#source_deletion_option}
	SourceDeletionOption *string `field:"optional" json:"sourceDeletionOption" yaml:"sourceDeletionOption"`
}

