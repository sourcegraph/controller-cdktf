package googlechroniclefeed


type GoogleChronicleFeedDetailsOffice365Settings struct {
	// API Auth Endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#auth_endpoint GoogleChronicleFeed#auth_endpoint}
	AuthEndpoint *string `field:"optional" json:"authEndpoint" yaml:"authEndpoint"`
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#authentication GoogleChronicleFeed#authentication}
	Authentication *GoogleChronicleFeedDetailsOffice365SettingsAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// Supported office 365 content type. Possible values: AUDIT_AZURE_ACTIVE_DIRECTORY AUDIT_EXCHANGE AUDIT_SHARE_POINT AUDIT_GENERAL DLP_ALL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#content_type GoogleChronicleFeed#content_type}
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// API Hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#hostname GoogleChronicleFeed#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Tenant ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#tenant_id GoogleChronicleFeed#tenant_id}
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

