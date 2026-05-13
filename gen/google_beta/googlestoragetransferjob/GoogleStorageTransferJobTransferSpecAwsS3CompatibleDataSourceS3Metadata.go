package googlestoragetransferjob


type GoogleStorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3Metadata struct {
	// Authentication and authorization method used by the storage service.
	//
	// When not specified, Transfer Service will attempt to determine right auth method to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_storage_transfer_job#auth_method GoogleStorageTransferJob#auth_method}
	AuthMethod *string `field:"optional" json:"authMethod" yaml:"authMethod"`
	// The Listing API to use for discovering objects.
	//
	// When not specified, Transfer Service will attempt to determine the right API to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_storage_transfer_job#list_api GoogleStorageTransferJob#list_api}
	ListApi *string `field:"optional" json:"listApi" yaml:"listApi"`
	// The network protocol of the agent. When not specified, the default value of NetworkProtocol NETWORK_PROTOCOL_HTTPS is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_storage_transfer_job#protocol GoogleStorageTransferJob#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// API request model used to call the storage service.
	//
	// When not specified, the default value of RequestModel REQUEST_MODEL_VIRTUAL_HOSTED_STYLE is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_storage_transfer_job#request_model GoogleStorageTransferJob#request_model}
	RequestModel *string `field:"optional" json:"requestModel" yaml:"requestModel"`
}

