package googleintegrationconnectorsconnection


type GoogleIntegrationConnectorsConnectionAuthConfigOauth2JwtBearer struct {
	// client_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#client_key GoogleIntegrationConnectorsConnection#client_key}
	ClientKey *GoogleIntegrationConnectorsConnectionAuthConfigOauth2JwtBearerClientKey `field:"optional" json:"clientKey" yaml:"clientKey"`
	// jwt_claims block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#jwt_claims GoogleIntegrationConnectorsConnection#jwt_claims}
	JwtClaims *GoogleIntegrationConnectorsConnectionAuthConfigOauth2JwtBearerJwtClaims `field:"optional" json:"jwtClaims" yaml:"jwtClaims"`
}

