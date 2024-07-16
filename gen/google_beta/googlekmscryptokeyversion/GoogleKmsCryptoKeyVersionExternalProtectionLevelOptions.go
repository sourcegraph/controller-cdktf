package googlekmscryptokeyversion


type GoogleKmsCryptoKeyVersionExternalProtectionLevelOptions struct {
	// The path to the external key material on the EKM when using EkmConnection e.g., "v0/my/key". Set this field instead of externalKeyUri when using an EkmConnection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_kms_crypto_key_version#ekm_connection_key_path GoogleKmsCryptoKeyVersion#ekm_connection_key_path}
	EkmConnectionKeyPath *string `field:"optional" json:"ekmConnectionKeyPath" yaml:"ekmConnectionKeyPath"`
	// The URI for an external resource that this CryptoKeyVersion represents.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_kms_crypto_key_version#external_key_uri GoogleKmsCryptoKeyVersion#external_key_uri}
	ExternalKeyUri *string `field:"optional" json:"externalKeyUri" yaml:"externalKeyUri"`
}

