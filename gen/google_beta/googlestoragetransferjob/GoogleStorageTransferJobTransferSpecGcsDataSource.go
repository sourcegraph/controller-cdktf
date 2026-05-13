package googlestoragetransferjob


type GoogleStorageTransferJobTransferSpecGcsDataSource struct {
	// Google Cloud Storage bucket name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_storage_transfer_job#bucket_name GoogleStorageTransferJob#bucket_name}
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Google Cloud Storage path in bucket to transfer.
	//
	// Must be an empty string or full path name that ends with a '/'. This field is treated as an object prefix. As such, it should not begin with a '/'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_storage_transfer_job#path GoogleStorageTransferJob#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

