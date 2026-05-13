package googlechroniclefeed


type GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthentication struct {
	// aws_iam_role_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#aws_iam_role_auth GoogleChronicleFeed#aws_iam_role_auth}
	AwsIamRoleAuth *GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationAwsIamRoleAuth `field:"required" json:"awsIamRoleAuth" yaml:"awsIamRoleAuth"`
	// sqs_v2_access_key_secret_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#sqs_v2_access_key_secret_auth GoogleChronicleFeed#sqs_v2_access_key_secret_auth}
	SqsV2AccessKeySecretAuth *GoogleChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationSqsV2AccessKeySecretAuth `field:"required" json:"sqsV2AccessKeySecretAuth" yaml:"sqsV2AccessKeySecretAuth"`
}

