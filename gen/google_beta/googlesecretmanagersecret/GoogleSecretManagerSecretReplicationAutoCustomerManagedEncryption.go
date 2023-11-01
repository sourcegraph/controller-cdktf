package googlesecretmanagersecret


type GoogleSecretManagerSecretReplicationAutoCustomerManagedEncryption struct {
	// The resource name of the Cloud KMS CryptoKey used to encrypt secret payloads.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.4.0/docs/resources/google_secret_manager_secret#kms_key_name GoogleSecretManagerSecret#kms_key_name}
	KmsKeyName *string `field:"required" json:"kmsKeyName" yaml:"kmsKeyName"`
}

