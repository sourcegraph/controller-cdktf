package storagetransferjob


type StorageTransferJobTransferSpecHdfsDataSource struct {
	// Directory path to the filesystem.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/storage_transfer_job#path StorageTransferJob#path}
	Path *string `field:"required" json:"path" yaml:"path"`
}

