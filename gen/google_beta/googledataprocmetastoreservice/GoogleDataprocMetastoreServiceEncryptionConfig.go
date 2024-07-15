package googledataprocmetastoreservice


type GoogleDataprocMetastoreServiceEncryptionConfig struct {
	// The fully qualified customer provided Cloud KMS key name to use for customer data encryption. Use the following format: 'projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_metastore_service#kms_key GoogleDataprocMetastoreService#kms_key}
	KmsKey *string `field:"required" json:"kmsKey" yaml:"kmsKey"`
}

