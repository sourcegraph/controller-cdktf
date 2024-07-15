package googledatastreamconnectionprofile


type GoogleDatastreamConnectionProfileMysqlProfile struct {
	// Hostname for the MySQL connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_connection_profile#hostname GoogleDatastreamConnectionProfile#hostname}
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// Password for the MySQL connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_connection_profile#password GoogleDatastreamConnectionProfile#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// Username for the MySQL connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_connection_profile#username GoogleDatastreamConnectionProfile#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// Port for the MySQL connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_connection_profile#port GoogleDatastreamConnectionProfile#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// ssl_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_connection_profile#ssl_config GoogleDatastreamConnectionProfile#ssl_config}
	SslConfig *GoogleDatastreamConnectionProfileMysqlProfileSslConfig `field:"optional" json:"sslConfig" yaml:"sslConfig"`
}

