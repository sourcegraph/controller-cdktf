package googlealloydbinstance


type GoogleAlloydbInstanceClientConnectionConfig struct {
	// Configuration to enforce connectors only (ex: AuthProxy) connections to the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_instance#require_connectors GoogleAlloydbInstance#require_connectors}
	RequireConnectors interface{} `field:"optional" json:"requireConnectors" yaml:"requireConnectors"`
	// ssl_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_instance#ssl_config GoogleAlloydbInstance#ssl_config}
	SslConfig *GoogleAlloydbInstanceClientConnectionConfigSslConfig `field:"optional" json:"sslConfig" yaml:"sslConfig"`
}

