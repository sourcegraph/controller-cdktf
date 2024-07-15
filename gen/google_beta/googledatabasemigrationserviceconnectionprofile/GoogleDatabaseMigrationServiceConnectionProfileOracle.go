package googledatabasemigrationserviceconnectionprofile


type GoogleDatabaseMigrationServiceConnectionProfileOracle struct {
	// Required. Database service for the Oracle connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_database_migration_service_connection_profile#database_service GoogleDatabaseMigrationServiceConnectionProfile#database_service}
	DatabaseService *string `field:"required" json:"databaseService" yaml:"databaseService"`
	// Required. The IP or hostname of the source Oracle database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_database_migration_service_connection_profile#host GoogleDatabaseMigrationServiceConnectionProfile#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// Required.
	//
	// Input only. The password for the user that Database Migration Service will be using to connect to the database.
	// This field is not returned on request, and the value is encrypted when stored in Database Migration Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_database_migration_service_connection_profile#password GoogleDatabaseMigrationServiceConnectionProfile#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// Required. The network port of the source Oracle database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_database_migration_service_connection_profile#port GoogleDatabaseMigrationServiceConnectionProfile#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Required.
	//
	// The username that Database Migration Service will use to connect to the database. The value is encrypted when stored in Database Migration Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_database_migration_service_connection_profile#username GoogleDatabaseMigrationServiceConnectionProfile#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// forward_ssh_connectivity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_database_migration_service_connection_profile#forward_ssh_connectivity GoogleDatabaseMigrationServiceConnectionProfile#forward_ssh_connectivity}
	ForwardSshConnectivity *GoogleDatabaseMigrationServiceConnectionProfileOracleForwardSshConnectivity `field:"optional" json:"forwardSshConnectivity" yaml:"forwardSshConnectivity"`
	// private_connectivity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_database_migration_service_connection_profile#private_connectivity GoogleDatabaseMigrationServiceConnectionProfile#private_connectivity}
	PrivateConnectivity *GoogleDatabaseMigrationServiceConnectionProfileOraclePrivateConnectivity `field:"optional" json:"privateConnectivity" yaml:"privateConnectivity"`
	// ssl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_database_migration_service_connection_profile#ssl GoogleDatabaseMigrationServiceConnectionProfile#ssl}
	Ssl *GoogleDatabaseMigrationServiceConnectionProfileOracleSsl `field:"optional" json:"ssl" yaml:"ssl"`
	// static_service_ip_connectivity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_database_migration_service_connection_profile#static_service_ip_connectivity GoogleDatabaseMigrationServiceConnectionProfile#static_service_ip_connectivity}
	StaticServiceIpConnectivity *GoogleDatabaseMigrationServiceConnectionProfileOracleStaticServiceIpConnectivity `field:"optional" json:"staticServiceIpConnectivity" yaml:"staticServiceIpConnectivity"`
}

