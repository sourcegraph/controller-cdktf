package googlestoragetransferjob


type GoogleStorageTransferJobTransferSpecPosixDataSink struct {
	// Root directory path to the filesystem.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_transfer_job#root_directory GoogleStorageTransferJob#root_directory}
	RootDirectory *string `field:"required" json:"rootDirectory" yaml:"rootDirectory"`
}

