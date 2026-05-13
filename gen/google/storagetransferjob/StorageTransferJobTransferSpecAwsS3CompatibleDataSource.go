package storagetransferjob


type StorageTransferJobTransferSpecAwsS3CompatibleDataSource struct {
	// Name of the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/storage_transfer_job#bucket_name StorageTransferJob#bucket_name}
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Endpoint of the storage service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/storage_transfer_job#endpoint StorageTransferJob#endpoint}
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// Specifies the path to transfer objects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/storage_transfer_job#path StorageTransferJob#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Specifies the region to sign requests with.
	//
	// This can be left blank if requests should be signed with an empty region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/storage_transfer_job#region StorageTransferJob#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// s3_metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/storage_transfer_job#s3_metadata StorageTransferJob#s3_metadata}
	S3Metadata *StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3Metadata `field:"optional" json:"s3Metadata" yaml:"s3Metadata"`
}

