package googledeveloperconnectaccountconnector


type GoogleDeveloperConnectAccountConnectorProviderOauthConfig struct {
	// User selected scopes to apply to the Oauth config In the event of changing scopes, user records under AccountConnector will be deleted and users will re-auth again.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_developer_connect_account_connector#scopes GoogleDeveloperConnectAccountConnector#scopes}
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// Possible values: GITHUB GITLAB GOOGLE SENTRY ROVO NEW_RELIC DATASTAX DYNATRACE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_developer_connect_account_connector#system_provider_id GoogleDeveloperConnectAccountConnector#system_provider_id}
	SystemProviderId *string `field:"optional" json:"systemProviderId" yaml:"systemProviderId"`
}

