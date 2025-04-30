package googlestoragetransferjob


type GoogleStorageTransferJobReplicationSpec struct {
	// gcs_data_sink block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_storage_transfer_job#gcs_data_sink GoogleStorageTransferJob#gcs_data_sink}
	GcsDataSink *GoogleStorageTransferJobReplicationSpecGcsDataSink `field:"optional" json:"gcsDataSink" yaml:"gcsDataSink"`
	// gcs_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_storage_transfer_job#gcs_data_source GoogleStorageTransferJob#gcs_data_source}
	GcsDataSource *GoogleStorageTransferJobReplicationSpecGcsDataSource `field:"optional" json:"gcsDataSource" yaml:"gcsDataSource"`
	// object_conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_storage_transfer_job#object_conditions GoogleStorageTransferJob#object_conditions}
	ObjectConditions *GoogleStorageTransferJobReplicationSpecObjectConditions `field:"optional" json:"objectConditions" yaml:"objectConditions"`
	// transfer_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_storage_transfer_job#transfer_options GoogleStorageTransferJob#transfer_options}
	TransferOptions *GoogleStorageTransferJobReplicationSpecTransferOptions `field:"optional" json:"transferOptions" yaml:"transferOptions"`
}

