package googleintegrationconnectorsconnection


type GoogleIntegrationConnectorsConnectionAuthConfigOauth2JwtBearerJwtClaims struct {
	// Value for the "aud" claim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#audience GoogleIntegrationConnectorsConnection#audience}
	Audience *string `field:"optional" json:"audience" yaml:"audience"`
	// Value for the "iss" claim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#issuer GoogleIntegrationConnectorsConnection#issuer}
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// Value for the "sub" claim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#subject GoogleIntegrationConnectorsConnection#subject}
	Subject *string `field:"optional" json:"subject" yaml:"subject"`
}

