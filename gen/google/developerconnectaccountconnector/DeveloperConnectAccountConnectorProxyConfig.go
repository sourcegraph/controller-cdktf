package developerconnectaccountconnector


type DeveloperConnectAccountConnectorProxyConfig struct {
	// Setting this to true allows the git and http proxies to perform actions on behalf of the user configured under the account connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/developer_connect_account_connector#enabled DeveloperConnectAccountConnector#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

