package googlestoragetransferjob


type GoogleStorageTransferJobTransferSpecAzureBlobStorageDataSourceFederatedIdentityConfig struct {
	// The client (application) ID of the application with federated credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_storage_transfer_job#client_id GoogleStorageTransferJob#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The tenant (directory) ID of the application with federated credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_storage_transfer_job#tenant_id GoogleStorageTransferJob#tenant_id}
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
}

