package datastreamconnectionprofile


type DatastreamConnectionProfileMongodbProfile struct {
	// host_addresses block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_connection_profile#host_addresses DatastreamConnectionProfile#host_addresses}
	HostAddresses interface{} `field:"required" json:"hostAddresses" yaml:"hostAddresses"`
	// Username for the MongoDB connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_connection_profile#username DatastreamConnectionProfile#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// Password for the MongoDB connection. Mutually exclusive with secretManagerStoredPassword.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_connection_profile#password DatastreamConnectionProfile#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Name of the replica set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_connection_profile#replica_set DatastreamConnectionProfile#replica_set}
	ReplicaSet *string `field:"optional" json:"replicaSet" yaml:"replicaSet"`
	// A reference to a Secret Manager resource name storing the MongoDB connection password. Mutually exclusive with password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_connection_profile#secret_manager_stored_password DatastreamConnectionProfile#secret_manager_stored_password}
	SecretManagerStoredPassword *string `field:"optional" json:"secretManagerStoredPassword" yaml:"secretManagerStoredPassword"`
	// srv_connection_format block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_connection_profile#srv_connection_format DatastreamConnectionProfile#srv_connection_format}
	SrvConnectionFormat *DatastreamConnectionProfileMongodbProfileSrvConnectionFormat `field:"optional" json:"srvConnectionFormat" yaml:"srvConnectionFormat"`
	// ssl_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_connection_profile#ssl_config DatastreamConnectionProfile#ssl_config}
	SslConfig *DatastreamConnectionProfileMongodbProfileSslConfig `field:"optional" json:"sslConfig" yaml:"sslConfig"`
	// standard_connection_format block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_connection_profile#standard_connection_format DatastreamConnectionProfile#standard_connection_format}
	StandardConnectionFormat *DatastreamConnectionProfileMongodbProfileStandardConnectionFormat `field:"optional" json:"standardConnectionFormat" yaml:"standardConnectionFormat"`
}

