package googledatabasemigrationserviceconnectionprofile


type GoogleDatabaseMigrationServiceConnectionProfilePostgresql struct {
	// If the connected database is an AlloyDB instance, use this field to provide the AlloyDB cluster ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_database_migration_service_connection_profile#alloydb_cluster_id GoogleDatabaseMigrationServiceConnectionProfile#alloydb_cluster_id}
	AlloydbClusterId *string `field:"optional" json:"alloydbClusterId" yaml:"alloydbClusterId"`
	// If the source is a Cloud SQL database, use this field to provide the Cloud SQL instance ID of the source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_database_migration_service_connection_profile#cloud_sql_id GoogleDatabaseMigrationServiceConnectionProfile#cloud_sql_id}
	CloudSqlId *string `field:"optional" json:"cloudSqlId" yaml:"cloudSqlId"`
	// The IP or hostname of the source MySQL database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_database_migration_service_connection_profile#host GoogleDatabaseMigrationServiceConnectionProfile#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Input only.
	//
	// The password for the user that Database Migration Service will be using to connect to the database.
	// This field is not returned on request, and the value is encrypted when stored in Database Migration Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_database_migration_service_connection_profile#password GoogleDatabaseMigrationServiceConnectionProfile#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The network port of the source MySQL database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_database_migration_service_connection_profile#port GoogleDatabaseMigrationServiceConnectionProfile#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// ssl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_database_migration_service_connection_profile#ssl GoogleDatabaseMigrationServiceConnectionProfile#ssl}
	Ssl *GoogleDatabaseMigrationServiceConnectionProfilePostgresqlSsl `field:"optional" json:"ssl" yaml:"ssl"`
	// The username that Database Migration Service will use to connect to the database.
	//
	// The value is encrypted when stored in Database Migration Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_database_migration_service_connection_profile#username GoogleDatabaseMigrationServiceConnectionProfile#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

