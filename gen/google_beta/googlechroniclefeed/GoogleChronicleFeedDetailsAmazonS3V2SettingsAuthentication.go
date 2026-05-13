package googlechroniclefeed


type GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthentication struct {
	// access_key_secret_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#access_key_secret_auth GoogleChronicleFeed#access_key_secret_auth}
	AccessKeySecretAuth *GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAccessKeySecretAuth `field:"optional" json:"accessKeySecretAuth" yaml:"accessKeySecretAuth"`
	// aws_iam_role_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#aws_iam_role_auth GoogleChronicleFeed#aws_iam_role_auth}
	AwsIamRoleAuth *GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuth `field:"optional" json:"awsIamRoleAuth" yaml:"awsIamRoleAuth"`
}

