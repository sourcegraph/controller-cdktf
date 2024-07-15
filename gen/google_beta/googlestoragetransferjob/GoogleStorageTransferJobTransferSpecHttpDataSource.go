package googlestoragetransferjob


type GoogleStorageTransferJobTransferSpecHttpDataSource struct {
	// The URL that points to the file that stores the object list entries.
	//
	// This file must allow public access. Currently, only URLs with HTTP and HTTPS schemes are supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_transfer_job#list_url GoogleStorageTransferJob#list_url}
	ListUrl *string `field:"required" json:"listUrl" yaml:"listUrl"`
}

