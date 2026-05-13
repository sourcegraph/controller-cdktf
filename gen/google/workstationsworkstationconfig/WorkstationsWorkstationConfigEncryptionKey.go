package workstationsworkstationconfig


type WorkstationsWorkstationConfigEncryptionKey struct {
	// The name of the Google Cloud KMS encryption key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config#kms_key WorkstationsWorkstationConfigA#kms_key}
	KmsKey *string `field:"required" json:"kmsKey" yaml:"kmsKey"`
	// The service account to use with the specified KMS key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config#kms_key_service_account WorkstationsWorkstationConfigA#kms_key_service_account}
	KmsKeyServiceAccount *string `field:"required" json:"kmsKeyServiceAccount" yaml:"kmsKeyServiceAccount"`
}

