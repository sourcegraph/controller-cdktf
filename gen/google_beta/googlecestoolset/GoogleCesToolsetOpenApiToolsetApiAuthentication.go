package googlecestoolset


type GoogleCesToolsetOpenApiToolsetApiAuthentication struct {
	// api_key_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_toolset#api_key_config GoogleCesToolset#api_key_config}
	ApiKeyConfig *GoogleCesToolsetOpenApiToolsetApiAuthenticationApiKeyConfig `field:"optional" json:"apiKeyConfig" yaml:"apiKeyConfig"`
	// bearer_token_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_toolset#bearer_token_config GoogleCesToolset#bearer_token_config}
	BearerTokenConfig *GoogleCesToolsetOpenApiToolsetApiAuthenticationBearerTokenConfig `field:"optional" json:"bearerTokenConfig" yaml:"bearerTokenConfig"`
	// oauth_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_toolset#oauth_config GoogleCesToolset#oauth_config}
	OauthConfig *GoogleCesToolsetOpenApiToolsetApiAuthenticationOauthConfig `field:"optional" json:"oauthConfig" yaml:"oauthConfig"`
	// service_account_auth_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_toolset#service_account_auth_config GoogleCesToolset#service_account_auth_config}
	ServiceAccountAuthConfig *GoogleCesToolsetOpenApiToolsetApiAuthenticationServiceAccountAuthConfig `field:"optional" json:"serviceAccountAuthConfig" yaml:"serviceAccountAuthConfig"`
	// service_agent_id_token_auth_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_toolset#service_agent_id_token_auth_config GoogleCesToolset#service_agent_id_token_auth_config}
	ServiceAgentIdTokenAuthConfig *GoogleCesToolsetOpenApiToolsetApiAuthenticationServiceAgentIdTokenAuthConfig `field:"optional" json:"serviceAgentIdTokenAuthConfig" yaml:"serviceAgentIdTokenAuthConfig"`
}

