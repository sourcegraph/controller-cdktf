package googlestoragetransferjob


type GoogleStorageTransferJobTransferSpecAwsS3DataSourceAwsAccessKey struct {
	// AWS Key ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_transfer_job#access_key_id GoogleStorageTransferJob#access_key_id}
	AccessKeyId *string `field:"required" json:"accessKeyId" yaml:"accessKeyId"`
	// AWS Secret Access Key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_transfer_job#secret_access_key GoogleStorageTransferJob#secret_access_key}
	SecretAccessKey *string `field:"required" json:"secretAccessKey" yaml:"secretAccessKey"`
}

