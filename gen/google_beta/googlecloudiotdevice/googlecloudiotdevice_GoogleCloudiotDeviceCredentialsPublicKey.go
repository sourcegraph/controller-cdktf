package googlecloudiotdevice


type GoogleCloudiotDeviceCredentialsPublicKey struct {
	// The format of the key. Possible values: ["RSA_PEM", "RSA_X509_PEM", "ES256_PEM", "ES256_X509_PEM"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloudiot_device#format GoogleCloudiotDevice#format}
	Format *string `field:"required" json:"format" yaml:"format"`
	// The key data.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloudiot_device#key GoogleCloudiotDevice#key}
	Key *string `field:"required" json:"key" yaml:"key"`
}

