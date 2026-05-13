package developerconnectconnection


type DeveloperConnectConnectionHttpConfig struct {
	// The service provider's https endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/developer_connect_connection#host_uri DeveloperConnectConnection#host_uri}
	HostUri *string `field:"required" json:"hostUri" yaml:"hostUri"`
	// basic_authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/developer_connect_connection#basic_authentication DeveloperConnectConnection#basic_authentication}
	BasicAuthentication *DeveloperConnectConnectionHttpConfigBasicAuthentication `field:"optional" json:"basicAuthentication" yaml:"basicAuthentication"`
	// bearer_token_authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/developer_connect_connection#bearer_token_authentication DeveloperConnectConnection#bearer_token_authentication}
	BearerTokenAuthentication *DeveloperConnectConnectionHttpConfigBearerTokenAuthentication `field:"optional" json:"bearerTokenAuthentication" yaml:"bearerTokenAuthentication"`
	// service_directory_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/developer_connect_connection#service_directory_config DeveloperConnectConnection#service_directory_config}
	ServiceDirectoryConfig *DeveloperConnectConnectionHttpConfigServiceDirectoryConfig `field:"optional" json:"serviceDirectoryConfig" yaml:"serviceDirectoryConfig"`
	// The SSL certificate to use for requests to the HTTP service provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/developer_connect_connection#ssl_ca_certificate DeveloperConnectConnection#ssl_ca_certificate}
	SslCaCertificate *string `field:"optional" json:"sslCaCertificate" yaml:"sslCaCertificate"`
}

