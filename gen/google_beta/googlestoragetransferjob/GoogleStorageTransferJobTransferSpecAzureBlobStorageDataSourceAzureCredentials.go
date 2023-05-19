package googlestoragetransferjob


type GoogleStorageTransferJobTransferSpecAzureBlobStorageDataSourceAzureCredentials struct {
	// Azure shared access signature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_storage_transfer_job#sas_token GoogleStorageTransferJob#sas_token}
	SasToken *string `field:"required" json:"sasToken" yaml:"sasToken"`
}

