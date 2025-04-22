package googlecomputeinstancefrommachineimage


type GoogleComputeInstanceFromMachineImageSourceMachineImageEncryptionKey struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_instance_from_machine_image#kms_key_name GoogleComputeInstanceFromMachineImage#kms_key_name}.
	KmsKeyName *string `field:"optional" json:"kmsKeyName" yaml:"kmsKeyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_instance_from_machine_image#kms_key_service_account GoogleComputeInstanceFromMachineImage#kms_key_service_account}.
	KmsKeyServiceAccount *string `field:"optional" json:"kmsKeyServiceAccount" yaml:"kmsKeyServiceAccount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_instance_from_machine_image#raw_key GoogleComputeInstanceFromMachineImage#raw_key}.
	RawKey *string `field:"optional" json:"rawKey" yaml:"rawKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_instance_from_machine_image#rsa_encrypted_key GoogleComputeInstanceFromMachineImage#rsa_encrypted_key}.
	RsaEncryptedKey *string `field:"optional" json:"rsaEncryptedKey" yaml:"rsaEncryptedKey"`
}

