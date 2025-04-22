package googlebigquerydatatransferconfig


type GoogleBigqueryDataTransferConfigEncryptionConfiguration struct {
	// The name of the KMS key used for encrypting BigQuery data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_bigquery_data_transfer_config#kms_key_name GoogleBigqueryDataTransferConfig#kms_key_name}
	KmsKeyName *string `field:"required" json:"kmsKeyName" yaml:"kmsKeyName"`
}

