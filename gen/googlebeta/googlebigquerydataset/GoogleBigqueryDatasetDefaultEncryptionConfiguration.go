package googlebigquerydataset


type GoogleBigqueryDatasetDefaultEncryptionConfiguration struct {
	// Describes the Cloud KMS encryption key that will be used to protect destination BigQuery table.
	//
	// The BigQuery Service Account associated with your project requires
	// access to this encryption key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigquery_dataset#kms_key_name GoogleBigqueryDataset#kms_key_name}
	KmsKeyName *string `field:"required" json:"kmsKeyName" yaml:"kmsKeyName"`
}

