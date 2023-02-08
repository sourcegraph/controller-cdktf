package devicepostureintegration


type DevicePostureIntegrationConfigA struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/device_posture_integration#api_url DevicePostureIntegration#api_url}.
	ApiUrl *string `field:"optional" json:"apiUrl" yaml:"apiUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/device_posture_integration#auth_url DevicePostureIntegration#auth_url}.
	AuthUrl *string `field:"optional" json:"authUrl" yaml:"authUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/device_posture_integration#client_id DevicePostureIntegration#client_id}.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/device_posture_integration#client_secret DevicePostureIntegration#client_secret}.
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
}

