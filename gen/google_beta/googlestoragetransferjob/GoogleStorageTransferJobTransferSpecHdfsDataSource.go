package googlestoragetransferjob


type GoogleStorageTransferJobTransferSpecHdfsDataSource struct {
	// Directory path to the filesystem.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_storage_transfer_job#path GoogleStorageTransferJob#path}
	Path *string `field:"required" json:"path" yaml:"path"`
}

