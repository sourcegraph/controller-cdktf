package dialogflowcxtoolversion


type DialogflowCxToolVersionToolOpenApiSpecAuthentication struct {
	// api_key_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_cx_tool_version#api_key_config DialogflowCxToolVersion#api_key_config}
	ApiKeyConfig *DialogflowCxToolVersionToolOpenApiSpecAuthenticationApiKeyConfig `field:"optional" json:"apiKeyConfig" yaml:"apiKeyConfig"`
	// bearer_token_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_cx_tool_version#bearer_token_config DialogflowCxToolVersion#bearer_token_config}
	BearerTokenConfig *DialogflowCxToolVersionToolOpenApiSpecAuthenticationBearerTokenConfig `field:"optional" json:"bearerTokenConfig" yaml:"bearerTokenConfig"`
	// oauth_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_cx_tool_version#oauth_config DialogflowCxToolVersion#oauth_config}
	OauthConfig *DialogflowCxToolVersionToolOpenApiSpecAuthenticationOauthConfig `field:"optional" json:"oauthConfig" yaml:"oauthConfig"`
	// service_agent_auth_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_cx_tool_version#service_agent_auth_config DialogflowCxToolVersion#service_agent_auth_config}
	ServiceAgentAuthConfig *DialogflowCxToolVersionToolOpenApiSpecAuthenticationServiceAgentAuthConfig `field:"optional" json:"serviceAgentAuthConfig" yaml:"serviceAgentAuthConfig"`
}

