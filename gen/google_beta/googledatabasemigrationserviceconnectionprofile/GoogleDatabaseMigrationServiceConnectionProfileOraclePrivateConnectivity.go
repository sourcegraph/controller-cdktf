package googledatabasemigrationserviceconnectionprofile


type GoogleDatabaseMigrationServiceConnectionProfileOraclePrivateConnectivity struct {
	// Required. The resource name (URI) of the private connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_database_migration_service_connection_profile#private_connection GoogleDatabaseMigrationServiceConnectionProfile#private_connection}
	PrivateConnection *string `field:"required" json:"privateConnection" yaml:"privateConnection"`
}

