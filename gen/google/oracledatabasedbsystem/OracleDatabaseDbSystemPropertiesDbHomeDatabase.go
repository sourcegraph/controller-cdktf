package oracledatabasedbsystem


type OracleDatabaseDbSystemPropertiesDbHomeDatabase struct {
	// The password for the default ADMIN user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#admin_password OracleDatabaseDbSystem#admin_password}
	AdminPassword *string `field:"required" json:"adminPassword" yaml:"adminPassword"`
	// The database ID of the Database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#database_id OracleDatabaseDbSystem#database_id}
	DatabaseId *string `field:"required" json:"databaseId" yaml:"databaseId"`
	// The character set for the database. The default is AL32UTF8.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#character_set OracleDatabaseDbSystem#character_set}
	CharacterSet *string `field:"optional" json:"characterSet" yaml:"characterSet"`
	// The name of the DbHome resource associated with the Database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#db_home_name OracleDatabaseDbSystem#db_home_name}
	DbHomeName *string `field:"optional" json:"dbHomeName" yaml:"dbHomeName"`
	// The database name.
	//
	// The name must begin with an alphabetic character and can
	// contain a maximum of eight alphanumeric characters. Special characters are
	// not permitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#db_name OracleDatabaseDbSystem#db_name}
	DbName *string `field:"optional" json:"dbName" yaml:"dbName"`
	// The DB_UNIQUE_NAME of the Oracle Database being backed up.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#db_unique_name OracleDatabaseDbSystem#db_unique_name}
	DbUniqueName *string `field:"optional" json:"dbUniqueName" yaml:"dbUniqueName"`
	// The GCP Oracle zone where the Database is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#gcp_oracle_zone OracleDatabaseDbSystem#gcp_oracle_zone}
	GcpOracleZone *string `field:"optional" json:"gcpOracleZone" yaml:"gcpOracleZone"`
	// The national character set for the database. The default is AL16UTF16.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#ncharacter_set OracleDatabaseDbSystem#ncharacter_set}
	NcharacterSet *string `field:"optional" json:"ncharacterSet" yaml:"ncharacterSet"`
	// The ID of the pluggable database associated with Database. The ID must be unique within the project and location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#pluggable_database_id OracleDatabaseDbSystem#pluggable_database_id}
	PluggableDatabaseId *string `field:"optional" json:"pluggableDatabaseId" yaml:"pluggableDatabaseId"`
	// The pluggable dataabse associated with the Database.
	//
	// The name must begin with an alphabetic character and can contain a maximum of thirty alphanumeric characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#pluggable_database_name OracleDatabaseDbSystem#pluggable_database_name}
	PluggableDatabaseName *string `field:"optional" json:"pluggableDatabaseName" yaml:"pluggableDatabaseName"`
	// properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#properties OracleDatabaseDbSystem#properties}
	Properties *OracleDatabaseDbSystemPropertiesDbHomeDatabaseProperties `field:"optional" json:"properties" yaml:"properties"`
	// The TDE wallet password for the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#tde_wallet_password OracleDatabaseDbSystem#tde_wallet_password}
	TdeWalletPassword *string `field:"optional" json:"tdeWalletPassword" yaml:"tdeWalletPassword"`
}

