package storagetransferjob


type StorageTransferJobTransferSpecTransferManifest struct {
	// Cloud Storage path to the manifest CSV.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/storage_transfer_job#location StorageTransferJob#location}
	Location *string `field:"required" json:"location" yaml:"location"`
}

