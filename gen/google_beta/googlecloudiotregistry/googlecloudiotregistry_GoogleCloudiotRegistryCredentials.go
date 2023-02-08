package googlecloudiotregistry


type GoogleCloudiotRegistryCredentials struct {
	// A public key certificate format and data.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloudiot_registry#public_key_certificate GoogleCloudiotRegistry#public_key_certificate}
	PublicKeyCertificate *map[string]*string `field:"required" json:"publicKeyCertificate" yaml:"publicKeyCertificate"`
}

