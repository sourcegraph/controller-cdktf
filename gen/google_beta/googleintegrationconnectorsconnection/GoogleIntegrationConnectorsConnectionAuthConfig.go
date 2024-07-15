package googleintegrationconnectorsconnection


type GoogleIntegrationConnectorsConnectionAuthConfig struct {
	// authType of the Connection Possible values: ["USER_PASSWORD", "OAUTH2_JWT_BEARER", "OAUTH2_CLIENT_CREDENTIALS", "SSH_PUBLIC_KEY", "OAUTH2_AUTH_CODE_FLOW"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#auth_type GoogleIntegrationConnectorsConnection#auth_type}
	AuthType *string `field:"required" json:"authType" yaml:"authType"`
	// additional_variable block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#additional_variable GoogleIntegrationConnectorsConnection#additional_variable}
	AdditionalVariable interface{} `field:"optional" json:"additionalVariable" yaml:"additionalVariable"`
	// The type of authentication configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#auth_key GoogleIntegrationConnectorsConnection#auth_key}
	AuthKey *string `field:"optional" json:"authKey" yaml:"authKey"`
	// oauth2_auth_code_flow block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#oauth2_auth_code_flow GoogleIntegrationConnectorsConnection#oauth2_auth_code_flow}
	Oauth2AuthCodeFlow *GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlow `field:"optional" json:"oauth2AuthCodeFlow" yaml:"oauth2AuthCodeFlow"`
	// oauth2_client_credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#oauth2_client_credentials GoogleIntegrationConnectorsConnection#oauth2_client_credentials}
	Oauth2ClientCredentials *GoogleIntegrationConnectorsConnectionAuthConfigOauth2ClientCredentials `field:"optional" json:"oauth2ClientCredentials" yaml:"oauth2ClientCredentials"`
	// oauth2_jwt_bearer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#oauth2_jwt_bearer GoogleIntegrationConnectorsConnection#oauth2_jwt_bearer}
	Oauth2JwtBearer *GoogleIntegrationConnectorsConnectionAuthConfigOauth2JwtBearer `field:"optional" json:"oauth2JwtBearer" yaml:"oauth2JwtBearer"`
	// ssh_public_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#ssh_public_key GoogleIntegrationConnectorsConnection#ssh_public_key}
	SshPublicKey *GoogleIntegrationConnectorsConnectionAuthConfigSshPublicKey `field:"optional" json:"sshPublicKey" yaml:"sshPublicKey"`
	// user_password block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#user_password GoogleIntegrationConnectorsConnection#user_password}
	UserPassword *GoogleIntegrationConnectorsConnectionAuthConfigUserPassword `field:"optional" json:"userPassword" yaml:"userPassword"`
}

