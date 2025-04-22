package googledatabasemigrationservicemigrationjob


type GoogleDatabaseMigrationServiceMigrationJobPerformanceConfig struct {
	// Initial dump parallelism level. Possible values: ["MIN", "OPTIMAL", "MAX"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_database_migration_service_migration_job#dump_parallel_level GoogleDatabaseMigrationServiceMigrationJob#dump_parallel_level}
	DumpParallelLevel *string `field:"optional" json:"dumpParallelLevel" yaml:"dumpParallelLevel"`
}

