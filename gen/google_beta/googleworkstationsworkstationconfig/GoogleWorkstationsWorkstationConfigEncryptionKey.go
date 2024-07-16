package googleworkstationsworkstationconfig


type GoogleWorkstationsWorkstationConfigEncryptionKey struct {
	// The name of the Google Cloud KMS encryption key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_config#kms_key GoogleWorkstationsWorkstationConfigA#kms_key}
	KmsKey *string `field:"required" json:"kmsKey" yaml:"kmsKey"`
	// The service account to use with the specified KMS key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_config#kms_key_service_account GoogleWorkstationsWorkstationConfigA#kms_key_service_account}
	KmsKeyServiceAccount *string `field:"required" json:"kmsKeyServiceAccount" yaml:"kmsKeyServiceAccount"`
}

