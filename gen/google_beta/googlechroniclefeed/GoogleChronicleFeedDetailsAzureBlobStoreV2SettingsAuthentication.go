package googlechroniclefeed


type GoogleChronicleFeedDetailsAzureBlobStoreV2SettingsAuthentication struct {
	// Access Key also known as shared key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#access_key GoogleChronicleFeed#access_key}
	AccessKey *string `field:"required" json:"accessKey" yaml:"accessKey"`
	// azure_v2_workload_identity_federation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#azure_v2_workload_identity_federation GoogleChronicleFeed#azure_v2_workload_identity_federation}
	AzureV2WorkloadIdentityFederation *GoogleChronicleFeedDetailsAzureBlobStoreV2SettingsAuthenticationAzureV2WorkloadIdentityFederation `field:"required" json:"azureV2WorkloadIdentityFederation" yaml:"azureV2WorkloadIdentityFederation"`
	// SAS Token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#sas_token GoogleChronicleFeed#sas_token}
	SasToken *string `field:"required" json:"sasToken" yaml:"sasToken"`
}

