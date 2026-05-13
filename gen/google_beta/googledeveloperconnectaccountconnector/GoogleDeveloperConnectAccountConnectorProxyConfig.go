package googledeveloperconnectaccountconnector


type GoogleDeveloperConnectAccountConnectorProxyConfig struct {
	// Setting this to true allows the git and http proxies to perform actions on behalf of the user configured under the account connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_developer_connect_account_connector#enabled GoogleDeveloperConnectAccountConnector#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

