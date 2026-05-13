package googledeveloperconnectconnection


type GoogleDeveloperConnectConnectionHttpConfig struct {
	// The service provider's https endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_developer_connect_connection#host_uri GoogleDeveloperConnectConnection#host_uri}
	HostUri *string `field:"required" json:"hostUri" yaml:"hostUri"`
	// basic_authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_developer_connect_connection#basic_authentication GoogleDeveloperConnectConnection#basic_authentication}
	BasicAuthentication *GoogleDeveloperConnectConnectionHttpConfigBasicAuthentication `field:"optional" json:"basicAuthentication" yaml:"basicAuthentication"`
	// bearer_token_authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_developer_connect_connection#bearer_token_authentication GoogleDeveloperConnectConnection#bearer_token_authentication}
	BearerTokenAuthentication *GoogleDeveloperConnectConnectionHttpConfigBearerTokenAuthentication `field:"optional" json:"bearerTokenAuthentication" yaml:"bearerTokenAuthentication"`
	// service_directory_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_developer_connect_connection#service_directory_config GoogleDeveloperConnectConnection#service_directory_config}
	ServiceDirectoryConfig *GoogleDeveloperConnectConnectionHttpConfigServiceDirectoryConfig `field:"optional" json:"serviceDirectoryConfig" yaml:"serviceDirectoryConfig"`
	// The SSL certificate to use for requests to the HTTP service provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_developer_connect_connection#ssl_ca_certificate GoogleDeveloperConnectConnection#ssl_ca_certificate}
	SslCaCertificate *string `field:"optional" json:"sslCaCertificate" yaml:"sslCaCertificate"`
}

