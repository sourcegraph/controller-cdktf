package googledatastreamconnectionprofile


type GoogleDatastreamConnectionProfileSqlServerProfile struct {
	// Database for the SQL Server connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_datastream_connection_profile#database GoogleDatastreamConnectionProfile#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Hostname for the SQL Server connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_datastream_connection_profile#hostname GoogleDatastreamConnectionProfile#hostname}
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// Username for the SQL Server connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_datastream_connection_profile#username GoogleDatastreamConnectionProfile#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// Password for the SQL Server connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_datastream_connection_profile#password GoogleDatastreamConnectionProfile#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Port for the SQL Server connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_datastream_connection_profile#port GoogleDatastreamConnectionProfile#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// A reference to a Secret Manager resource name storing the user's password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_datastream_connection_profile#secret_manager_stored_password GoogleDatastreamConnectionProfile#secret_manager_stored_password}
	SecretManagerStoredPassword *string `field:"optional" json:"secretManagerStoredPassword" yaml:"secretManagerStoredPassword"`
}

