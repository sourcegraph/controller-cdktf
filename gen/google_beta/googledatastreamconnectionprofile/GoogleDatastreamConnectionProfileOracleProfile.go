package googledatastreamconnectionprofile


type GoogleDatastreamConnectionProfileOracleProfile struct {
	// Database for the Oracle connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_connection_profile#database_service GoogleDatastreamConnectionProfile#database_service}
	DatabaseService *string `field:"required" json:"databaseService" yaml:"databaseService"`
	// Hostname for the Oracle connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_connection_profile#hostname GoogleDatastreamConnectionProfile#hostname}
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// Password for the Oracle connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_connection_profile#password GoogleDatastreamConnectionProfile#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// Username for the Oracle connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_connection_profile#username GoogleDatastreamConnectionProfile#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// Connection string attributes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_connection_profile#connection_attributes GoogleDatastreamConnectionProfile#connection_attributes}
	ConnectionAttributes *map[string]*string `field:"optional" json:"connectionAttributes" yaml:"connectionAttributes"`
	// Port for the Oracle connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_connection_profile#port GoogleDatastreamConnectionProfile#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

