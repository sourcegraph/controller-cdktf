package storagetransferjob


type StorageTransferJobReplicationSpecGcsDataSource struct {
	// Google Cloud Storage bucket name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/storage_transfer_job#bucket_name StorageTransferJob#bucket_name}
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Google Cloud Storage path in bucket to transfer.
	//
	// Must be an empty string or full path name that ends with a '/'. This field is treated as an object prefix. As such, it should not begin with a '/'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/storage_transfer_job#path StorageTransferJob#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

