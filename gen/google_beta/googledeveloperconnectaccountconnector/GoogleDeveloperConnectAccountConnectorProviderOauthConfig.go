package googledeveloperconnectaccountconnector


type GoogleDeveloperConnectAccountConnectorProviderOauthConfig struct {
	// Required.
	//
	// User selected scopes to apply to the Oauth config
	// In the event of changing scopes, user records under AccountConnector will
	// be deleted and users will re-auth again.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_developer_connect_account_connector#scopes GoogleDeveloperConnectAccountConnector#scopes}
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// List of providers that are owned by Developer Connect.
	//
	// Possible values:
	// GITHUB
	// GITLAB
	// GOOGLE
	// SENTRY
	// ROVO
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_developer_connect_account_connector#system_provider_id GoogleDeveloperConnectAccountConnector#system_provider_id}
	SystemProviderId *string `field:"optional" json:"systemProviderId" yaml:"systemProviderId"`
}

