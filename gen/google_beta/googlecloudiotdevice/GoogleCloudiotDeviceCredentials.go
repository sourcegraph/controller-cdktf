package googlecloudiotdevice


type GoogleCloudiotDeviceCredentials struct {
	// public_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_cloudiot_device#public_key GoogleCloudiotDevice#public_key}
	PublicKey *GoogleCloudiotDeviceCredentialsPublicKey `field:"required" json:"publicKey" yaml:"publicKey"`
	// The time at which this credential becomes invalid.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_cloudiot_device#expiration_time GoogleCloudiotDevice#expiration_time}
	ExpirationTime *string `field:"optional" json:"expirationTime" yaml:"expirationTime"`
}

