package googleapigatewayapiconfig


type GoogleApiGatewayApiConfigGatewayConfigBackendConfig struct {
	// Google Cloud IAM service account used to sign OIDC tokens for backends that have authentication configured (https://cloud.google.com/service-infrastructure/docs/service-management/reference/rest/v1/services.configs#backend).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_api_gateway_api_config#google_service_account GoogleApiGatewayApiConfigA#google_service_account}
	GoogleServiceAccount *string `field:"required" json:"googleServiceAccount" yaml:"googleServiceAccount"`
}

