package googledeveloperconnectaccountconnector


type GoogleDeveloperConnectAccountConnectorCustomOauthConfig struct {
	// The OAuth2 authrization server URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_developer_connect_account_connector#auth_uri GoogleDeveloperConnectAccountConnector#auth_uri}
	AuthUri *string `field:"required" json:"authUri" yaml:"authUri"`
	// The client ID of the OAuth application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_developer_connect_account_connector#client_id GoogleDeveloperConnectAccountConnector#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Input only.
	//
	// The client secret of the OAuth application.
	// It will be provided as plain text, but encrypted and stored in developer
	// connect. As INPUT_ONLY field, it will not be included in the output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_developer_connect_account_connector#client_secret GoogleDeveloperConnectAccountConnector#client_secret}
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// The host URI of the OAuth application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_developer_connect_account_connector#host_uri GoogleDeveloperConnectAccountConnector#host_uri}
	HostUri *string `field:"required" json:"hostUri" yaml:"hostUri"`
	// The type of the SCM provider. Possible values: SCM_PROVIDER_UNKNOWN GITHUB_ENTERPRISE GITLAB_ENTERPRISE BITBUCKET_DATA_CENTER.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_developer_connect_account_connector#scm_provider GoogleDeveloperConnectAccountConnector#scm_provider}
	ScmProvider *string `field:"required" json:"scmProvider" yaml:"scmProvider"`
	// The scopes to be requested during OAuth.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_developer_connect_account_connector#scopes GoogleDeveloperConnectAccountConnector#scopes}
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// The OAuth2 token request URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_developer_connect_account_connector#token_uri GoogleDeveloperConnectAccountConnector#token_uri}
	TokenUri *string `field:"required" json:"tokenUri" yaml:"tokenUri"`
	// Disable PKCE for this OAuth config. PKCE is enabled by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_developer_connect_account_connector#pkce_disabled GoogleDeveloperConnectAccountConnector#pkce_disabled}
	PkceDisabled interface{} `field:"optional" json:"pkceDisabled" yaml:"pkceDisabled"`
	// service_directory_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_developer_connect_account_connector#service_directory_config GoogleDeveloperConnectAccountConnector#service_directory_config}
	ServiceDirectoryConfig *GoogleDeveloperConnectAccountConnectorCustomOauthConfigServiceDirectoryConfig `field:"optional" json:"serviceDirectoryConfig" yaml:"serviceDirectoryConfig"`
	// SSL certificate to use for requests to a private service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_developer_connect_account_connector#ssl_ca_certificate GoogleDeveloperConnectAccountConnector#ssl_ca_certificate}
	SslCaCertificate *string `field:"optional" json:"sslCaCertificate" yaml:"sslCaCertificate"`
}

