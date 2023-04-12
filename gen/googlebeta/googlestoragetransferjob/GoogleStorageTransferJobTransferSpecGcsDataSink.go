package googlestoragetransferjob


type GoogleStorageTransferJobTransferSpecGcsDataSink struct {
	// Google Cloud Storage bucket name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_transfer_job#bucket_name GoogleStorageTransferJob#bucket_name}
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Google Cloud Storage path in bucket to transfer.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_transfer_job#path GoogleStorageTransferJob#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

