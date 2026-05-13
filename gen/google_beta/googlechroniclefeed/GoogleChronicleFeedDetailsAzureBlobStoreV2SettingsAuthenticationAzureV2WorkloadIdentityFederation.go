package googlechroniclefeed


type GoogleChronicleFeedDetailsAzureBlobStoreV2SettingsAuthenticationAzureV2WorkloadIdentityFederation struct {
	// OAuth client ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#client_id GoogleChronicleFeed#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Subject ID of the Azure subscription.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#subject_id GoogleChronicleFeed#subject_id}
	SubjectId *string `field:"required" json:"subjectId" yaml:"subjectId"`
	// Tenant ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#tenant_id GoogleChronicleFeed#tenant_id}
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
}

