package oracledatabasedbsystem


type OracleDatabaseDbSystemPropertiesDbHome struct {
	// database block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#database OracleDatabaseDbSystem#database}
	Database *OracleDatabaseDbSystemPropertiesDbHomeDatabase `field:"required" json:"database" yaml:"database"`
	// A valid Oracle Database version. For a list of supported versions, use the ListDbVersions operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#db_version OracleDatabaseDbSystem#db_version}
	DbVersion *string `field:"required" json:"dbVersion" yaml:"dbVersion"`
	// The display name for the Database Home. The name does not have to be unique within your project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#display_name OracleDatabaseDbSystem#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Whether unified auditing is enabled for the Database Home.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#is_unified_auditing_enabled OracleDatabaseDbSystem#is_unified_auditing_enabled}
	IsUnifiedAuditingEnabled interface{} `field:"optional" json:"isUnifiedAuditingEnabled" yaml:"isUnifiedAuditingEnabled"`
}

