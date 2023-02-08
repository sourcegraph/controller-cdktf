package googlestoragetransferjob


type GoogleStorageTransferJobTransferSpecAzureBlobStorageDataSourceAzureCredentials struct {
	// Azure shared access signature.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_transfer_job#sas_token GoogleStorageTransferJob#sas_token}
	SasToken *string `field:"required" json:"sasToken" yaml:"sasToken"`
}

