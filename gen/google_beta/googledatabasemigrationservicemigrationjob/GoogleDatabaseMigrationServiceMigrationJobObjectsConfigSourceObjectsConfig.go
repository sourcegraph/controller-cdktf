package googledatabasemigrationservicemigrationjob


type GoogleDatabaseMigrationServiceMigrationJobObjectsConfigSourceObjectsConfig struct {
	// object_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_database_migration_service_migration_job#object_configs GoogleDatabaseMigrationServiceMigrationJob#object_configs}
	ObjectConfigs interface{} `field:"optional" json:"objectConfigs" yaml:"objectConfigs"`
	// The objects selection type of the migration job.
	//
	// When set to
	// 'SPECIFIED_OBJECTS', only the objects listed in 'objectConfigs' are
	// migrated. When set to 'ALL_OBJECTS', all objects available on the
	// source are migrated. Possible values: ["ALL_OBJECTS", "SPECIFIED_OBJECTS"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_database_migration_service_migration_job#objects_selection_type GoogleDatabaseMigrationServiceMigrationJob#objects_selection_type}
	ObjectsSelectionType *string `field:"optional" json:"objectsSelectionType" yaml:"objectsSelectionType"`
}

