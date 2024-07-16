package googlespannerdatabase


type GoogleSpannerDatabaseEncryptionConfig struct {
	// Fully qualified name of the KMS key to use to encrypt this database.
	//
	// This key must exist
	// in the same location as the Spanner Database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_spanner_database#kms_key_name GoogleSpannerDatabase#kms_key_name}
	KmsKeyName *string `field:"required" json:"kmsKeyName" yaml:"kmsKeyName"`
}

