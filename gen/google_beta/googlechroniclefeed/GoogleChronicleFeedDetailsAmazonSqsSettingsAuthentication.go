package googlechroniclefeed


type GoogleChronicleFeedDetailsAmazonSqsSettingsAuthentication struct {
	// additional_s3_access_key_secret_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#additional_s3_access_key_secret_auth GoogleChronicleFeed#additional_s3_access_key_secret_auth}
	AdditionalS3AccessKeySecretAuth *GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationAdditionalS3AccessKeySecretAuth `field:"optional" json:"additionalS3AccessKeySecretAuth" yaml:"additionalS3AccessKeySecretAuth"`
	// sqs_access_key_secret_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#sqs_access_key_secret_auth GoogleChronicleFeed#sqs_access_key_secret_auth}
	SqsAccessKeySecretAuth *GoogleChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuth `field:"optional" json:"sqsAccessKeySecretAuth" yaml:"sqsAccessKeySecretAuth"`
}

