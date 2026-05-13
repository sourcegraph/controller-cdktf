package databasemigrationservicemigrationjob


type DatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfigObjectConfigsObjectIdentifier struct {
	// The category of the migration job object: 'DATABASE', 'SCHEMA', or 'TABLE'. Possible values: ["DATABASE", "SCHEMA", "TABLE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/database_migration_service_migration_job#type DatabaseMigrationServiceMigrationJob#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The database name. Required only if the object uses a database name as part of its unique identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/database_migration_service_migration_job#database DatabaseMigrationServiceMigrationJob#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// The schema name. Required only if the object uses a schema name as part of its unique identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/database_migration_service_migration_job#schema DatabaseMigrationServiceMigrationJob#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
	// The table name. Required only if the object is a level below database or schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/database_migration_service_migration_job#table DatabaseMigrationServiceMigrationJob#table}
	Table *string `field:"optional" json:"table" yaml:"table"`
}

