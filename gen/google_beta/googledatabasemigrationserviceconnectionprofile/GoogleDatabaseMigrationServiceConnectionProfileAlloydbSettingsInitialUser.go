package googledatabasemigrationserviceconnectionprofile


type GoogleDatabaseMigrationServiceConnectionProfileAlloydbSettingsInitialUser struct {
	// The initial password for the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_database_migration_service_connection_profile#password GoogleDatabaseMigrationServiceConnectionProfile#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// The database username.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_database_migration_service_connection_profile#user GoogleDatabaseMigrationServiceConnectionProfile#user}
	User *string `field:"required" json:"user" yaml:"user"`
}

