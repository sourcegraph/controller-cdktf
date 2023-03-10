package connectinstancestorageconfig


type ConnectInstanceStorageConfigStorageConfigS3Config struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/connect_instance_storage_config#bucket_name ConnectInstanceStorageConfig#bucket_name}.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/connect_instance_storage_config#bucket_prefix ConnectInstanceStorageConfig#bucket_prefix}.
	BucketPrefix *string `field:"required" json:"bucketPrefix" yaml:"bucketPrefix"`
	// encryption_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/connect_instance_storage_config#encryption_config ConnectInstanceStorageConfig#encryption_config}
	EncryptionConfig *ConnectInstanceStorageConfigStorageConfigS3ConfigEncryptionConfig `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
}

