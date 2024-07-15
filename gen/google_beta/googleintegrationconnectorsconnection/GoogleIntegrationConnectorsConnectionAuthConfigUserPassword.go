package googleintegrationconnectorsconnection


type GoogleIntegrationConnectorsConnectionAuthConfigUserPassword struct {
	// Username for Authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#username GoogleIntegrationConnectorsConnection#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// password block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#password GoogleIntegrationConnectorsConnection#password}
	Password *GoogleIntegrationConnectorsConnectionAuthConfigUserPasswordPassword `field:"optional" json:"password" yaml:"password"`
}

