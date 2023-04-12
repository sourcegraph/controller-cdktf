package googleapigatewayapiconfig


type GoogleApiGatewayApiConfigGatewayConfig struct {
	// backend_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_api_gateway_api_config#backend_config GoogleApiGatewayApiConfigA#backend_config}
	BackendConfig *GoogleApiGatewayApiConfigGatewayConfigBackendConfig `field:"required" json:"backendConfig" yaml:"backendConfig"`
}

