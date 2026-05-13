package cestoolset


type CesToolsetOpenApiToolsetApiAuthenticationOauthConfig struct {
	// The client ID from the OAuth provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_toolset#client_id CesToolset#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The name of the SecretManager secret version resource storing the client secret. Format: 'projects/{project}/secrets/{secret}/versions/{version}'.
	//
	// Note: You should grant 'roles/secretmanager.secretAccessor' role to the CES
	// service agent
	// 'service-@gcp-sa-ces.iam.gserviceaccount.com'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_toolset#client_secret_version CesToolset#client_secret_version}
	ClientSecretVersion *string `field:"required" json:"clientSecretVersion" yaml:"clientSecretVersion"`
	// OAuth grant types. Possible values: CLIENT_CREDENTIAL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_toolset#oauth_grant_type CesToolset#oauth_grant_type}
	OauthGrantType *string `field:"required" json:"oauthGrantType" yaml:"oauthGrantType"`
	// The token endpoint in the OAuth provider to exchange for an access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_toolset#token_endpoint CesToolset#token_endpoint}
	TokenEndpoint *string `field:"required" json:"tokenEndpoint" yaml:"tokenEndpoint"`
	// The OAuth scopes to grant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_toolset#scopes CesToolset#scopes}
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
}

