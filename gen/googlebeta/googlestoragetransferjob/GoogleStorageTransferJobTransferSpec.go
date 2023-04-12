package googlestoragetransferjob


type GoogleStorageTransferJobTransferSpec struct {
	// aws_s3_data_source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_transfer_job#aws_s3_data_source GoogleStorageTransferJob#aws_s3_data_source}
	AwsS3DataSource *GoogleStorageTransferJobTransferSpecAwsS3DataSource `field:"optional" json:"awsS3DataSource" yaml:"awsS3DataSource"`
	// azure_blob_storage_data_source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_transfer_job#azure_blob_storage_data_source GoogleStorageTransferJob#azure_blob_storage_data_source}
	AzureBlobStorageDataSource *GoogleStorageTransferJobTransferSpecAzureBlobStorageDataSource `field:"optional" json:"azureBlobStorageDataSource" yaml:"azureBlobStorageDataSource"`
	// gcs_data_sink block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_transfer_job#gcs_data_sink GoogleStorageTransferJob#gcs_data_sink}
	GcsDataSink *GoogleStorageTransferJobTransferSpecGcsDataSink `field:"optional" json:"gcsDataSink" yaml:"gcsDataSink"`
	// gcs_data_source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_transfer_job#gcs_data_source GoogleStorageTransferJob#gcs_data_source}
	GcsDataSource *GoogleStorageTransferJobTransferSpecGcsDataSource `field:"optional" json:"gcsDataSource" yaml:"gcsDataSource"`
	// http_data_source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_transfer_job#http_data_source GoogleStorageTransferJob#http_data_source}
	HttpDataSource *GoogleStorageTransferJobTransferSpecHttpDataSource `field:"optional" json:"httpDataSource" yaml:"httpDataSource"`
	// object_conditions block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_transfer_job#object_conditions GoogleStorageTransferJob#object_conditions}
	ObjectConditions *GoogleStorageTransferJobTransferSpecObjectConditions `field:"optional" json:"objectConditions" yaml:"objectConditions"`
	// posix_data_sink block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_transfer_job#posix_data_sink GoogleStorageTransferJob#posix_data_sink}
	PosixDataSink *GoogleStorageTransferJobTransferSpecPosixDataSink `field:"optional" json:"posixDataSink" yaml:"posixDataSink"`
	// posix_data_source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_transfer_job#posix_data_source GoogleStorageTransferJob#posix_data_source}
	PosixDataSource *GoogleStorageTransferJobTransferSpecPosixDataSource `field:"optional" json:"posixDataSource" yaml:"posixDataSource"`
	// transfer_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_transfer_job#transfer_options GoogleStorageTransferJob#transfer_options}
	TransferOptions *GoogleStorageTransferJobTransferSpecTransferOptions `field:"optional" json:"transferOptions" yaml:"transferOptions"`
}

