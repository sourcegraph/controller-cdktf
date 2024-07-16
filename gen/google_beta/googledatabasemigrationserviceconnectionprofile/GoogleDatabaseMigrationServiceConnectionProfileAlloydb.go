package googledatabasemigrationserviceconnectionprofile


type GoogleDatabaseMigrationServiceConnectionProfileAlloydb struct {
	// Required. The AlloyDB cluster ID that this connection profile is associated with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_database_migration_service_connection_profile#cluster_id GoogleDatabaseMigrationServiceConnectionProfile#cluster_id}
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_database_migration_service_connection_profile#settings GoogleDatabaseMigrationServiceConnectionProfile#settings}
	Settings *GoogleDatabaseMigrationServiceConnectionProfileAlloydbSettings `field:"optional" json:"settings" yaml:"settings"`
}

