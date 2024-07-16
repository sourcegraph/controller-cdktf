package googlealloydbinstance


type GoogleAlloydbInstanceClientConnectionConfigSslConfig struct {
	// SSL mode. Specifies client-server SSL/TLS connection behavior. Possible values: ["ENCRYPTED_ONLY", "ALLOW_UNENCRYPTED_AND_ENCRYPTED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_instance#ssl_mode GoogleAlloydbInstance#ssl_mode}
	SslMode *string `field:"optional" json:"sslMode" yaml:"sslMode"`
}

