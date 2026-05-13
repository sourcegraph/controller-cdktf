package googlechroniclefeed


type GoogleChronicleFeedDetailsAzureAdContextSettings struct {
	// API Auth Endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#auth_endpoint GoogleChronicleFeed#auth_endpoint}
	AuthEndpoint *string `field:"optional" json:"authEndpoint" yaml:"authEndpoint"`
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#authentication GoogleChronicleFeed#authentication}
	Authentication *GoogleChronicleFeedDetailsAzureAdContextSettingsAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// API Hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#hostname GoogleChronicleFeed#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Whether to retrieve device information in user context.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#retrieve_devices GoogleChronicleFeed#retrieve_devices}
	RetrieveDevices interface{} `field:"optional" json:"retrieveDevices" yaml:"retrieveDevices"`
	// Whether to retrieve group information in user context.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#retrieve_groups GoogleChronicleFeed#retrieve_groups}
	RetrieveGroups interface{} `field:"optional" json:"retrieveGroups" yaml:"retrieveGroups"`
	// Tenant ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#tenant_id GoogleChronicleFeed#tenant_id}
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

