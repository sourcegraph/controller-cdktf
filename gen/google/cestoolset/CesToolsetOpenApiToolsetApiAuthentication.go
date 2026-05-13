package cestoolset


type CesToolsetOpenApiToolsetApiAuthentication struct {
	// api_key_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_toolset#api_key_config CesToolset#api_key_config}
	ApiKeyConfig *CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfig `field:"optional" json:"apiKeyConfig" yaml:"apiKeyConfig"`
	// bearer_token_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_toolset#bearer_token_config CesToolset#bearer_token_config}
	BearerTokenConfig *CesToolsetOpenApiToolsetApiAuthenticationBearerTokenConfig `field:"optional" json:"bearerTokenConfig" yaml:"bearerTokenConfig"`
	// oauth_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_toolset#oauth_config CesToolset#oauth_config}
	OauthConfig *CesToolsetOpenApiToolsetApiAuthenticationOauthConfig `field:"optional" json:"oauthConfig" yaml:"oauthConfig"`
	// service_account_auth_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_toolset#service_account_auth_config CesToolset#service_account_auth_config}
	ServiceAccountAuthConfig *CesToolsetOpenApiToolsetApiAuthenticationServiceAccountAuthConfig `field:"optional" json:"serviceAccountAuthConfig" yaml:"serviceAccountAuthConfig"`
	// service_agent_id_token_auth_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_toolset#service_agent_id_token_auth_config CesToolset#service_agent_id_token_auth_config}
	ServiceAgentIdTokenAuthConfig *CesToolsetOpenApiToolsetApiAuthenticationServiceAgentIdTokenAuthConfig `field:"optional" json:"serviceAgentIdTokenAuthConfig" yaml:"serviceAgentIdTokenAuthConfig"`
}

