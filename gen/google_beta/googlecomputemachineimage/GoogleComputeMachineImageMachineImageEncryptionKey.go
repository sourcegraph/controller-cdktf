package googlecomputemachineimage


type GoogleComputeMachineImageMachineImageEncryptionKey struct {
	// The name of the encryption key that is stored in Google Cloud KMS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_machine_image#kms_key_name GoogleComputeMachineImage#kms_key_name}
	KmsKeyName *string `field:"optional" json:"kmsKeyName" yaml:"kmsKeyName"`
	// The service account used for the encryption request for the given KMS key.
	//
	// If absent, the Compute Engine Service Agent service account is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_machine_image#kms_key_service_account GoogleComputeMachineImage#kms_key_service_account}
	KmsKeyServiceAccount *string `field:"optional" json:"kmsKeyServiceAccount" yaml:"kmsKeyServiceAccount"`
	// Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to either encrypt or decrypt this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_machine_image#raw_key GoogleComputeMachineImage#raw_key}
	RawKey *string `field:"optional" json:"rawKey" yaml:"rawKey"`
}

