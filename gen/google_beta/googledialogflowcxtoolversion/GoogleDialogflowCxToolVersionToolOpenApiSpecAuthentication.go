package googledialogflowcxtoolversion


type GoogleDialogflowCxToolVersionToolOpenApiSpecAuthentication struct {
	// api_key_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool_version#api_key_config GoogleDialogflowCxToolVersion#api_key_config}
	ApiKeyConfig *GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationApiKeyConfig `field:"optional" json:"apiKeyConfig" yaml:"apiKeyConfig"`
	// bearer_token_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool_version#bearer_token_config GoogleDialogflowCxToolVersion#bearer_token_config}
	BearerTokenConfig *GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationBearerTokenConfig `field:"optional" json:"bearerTokenConfig" yaml:"bearerTokenConfig"`
	// oauth_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool_version#oauth_config GoogleDialogflowCxToolVersion#oauth_config}
	OauthConfig *GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOauthConfig `field:"optional" json:"oauthConfig" yaml:"oauthConfig"`
	// service_agent_auth_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool_version#service_agent_auth_config GoogleDialogflowCxToolVersion#service_agent_auth_config}
	ServiceAgentAuthConfig *GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationServiceAgentAuthConfig `field:"optional" json:"serviceAgentAuthConfig" yaml:"serviceAgentAuthConfig"`
}

