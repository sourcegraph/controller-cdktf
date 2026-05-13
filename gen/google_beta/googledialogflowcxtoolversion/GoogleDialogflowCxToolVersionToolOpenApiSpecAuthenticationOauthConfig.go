package googledialogflowcxtoolversion


type GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOauthConfig struct {
	// The client ID from the OAuth provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool_version#client_id GoogleDialogflowCxToolVersion#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// OAuth grant types. See [OauthGrantType](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.tools#oauthgranttype) for valid values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool_version#oauth_grant_type GoogleDialogflowCxToolVersion#oauth_grant_type}
	OauthGrantType *string `field:"required" json:"oauthGrantType" yaml:"oauthGrantType"`
	// The token endpoint in the OAuth provider to exchange for an access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool_version#token_endpoint GoogleDialogflowCxToolVersion#token_endpoint}
	TokenEndpoint *string `field:"required" json:"tokenEndpoint" yaml:"tokenEndpoint"`
	// Optional. The client secret from the OAuth provider. If the 'secretVersionForClientSecret' field is set, this field will be ignored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool_version#client_secret GoogleDialogflowCxToolVersion#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Optional. The OAuth scopes to grant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool_version#scopes GoogleDialogflowCxToolVersion#scopes}
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
	// Optional.
	//
	// The name of the SecretManager secret version resource storing the client secret.
	// If this field is set, the clientSecret field will be ignored.
	// Format: projects/{project}/secrets/{secret}/versions/{version}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool_version#secret_version_for_client_secret GoogleDialogflowCxToolVersion#secret_version_for_client_secret}
	SecretVersionForClientSecret *string `field:"optional" json:"secretVersionForClientSecret" yaml:"secretVersionForClientSecret"`
}

