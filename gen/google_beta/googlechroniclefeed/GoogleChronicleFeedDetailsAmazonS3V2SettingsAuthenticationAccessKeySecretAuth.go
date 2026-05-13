package googlechroniclefeed


type GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAccessKeySecretAuth struct {
	// Access Key ID for an AWS account (a 20-character, alphanumeric string).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#access_key_id GoogleChronicleFeed#access_key_id}
	AccessKeyId *string `field:"required" json:"accessKeyId" yaml:"accessKeyId"`
	// Secret Access Key for an AWS account (a 40-character string).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#secret_access_key GoogleChronicleFeed#secret_access_key}
	SecretAccessKey *string `field:"required" json:"secretAccessKey" yaml:"secretAccessKey"`
}

