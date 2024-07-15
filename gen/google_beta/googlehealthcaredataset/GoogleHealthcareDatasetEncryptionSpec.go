package googlehealthcaredataset


type GoogleHealthcareDatasetEncryptionSpec struct {
	// KMS encryption key that is used to secure this dataset and its sub-resources.
	//
	// The key used for
	// encryption and the dataset must be in the same location. If empty, the default Google encryption
	// key will be used to secure this dataset. The format is
	// projects/{projectId}/locations/{locationId}/keyRings/{keyRingId}/cryptoKeys/{keyId}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_healthcare_dataset#kms_key_name GoogleHealthcareDataset#kms_key_name}
	KmsKeyName *string `field:"optional" json:"kmsKeyName" yaml:"kmsKeyName"`
}

