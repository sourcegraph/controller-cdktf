package googlealloydbcluster


type GoogleAlloydbClusterEncryptionConfig struct {
	// The fully-qualified resource name of the KMS key.
	//
	// Each Cloud KMS key is regionalized and has the following format: projects/[PROJECT]/locations/[REGION]/keyRings/[RING]/cryptoKeys/[KEY_NAME].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#kms_key_name GoogleAlloydbCluster#kms_key_name}
	KmsKeyName *string `field:"optional" json:"kmsKeyName" yaml:"kmsKeyName"`
}

