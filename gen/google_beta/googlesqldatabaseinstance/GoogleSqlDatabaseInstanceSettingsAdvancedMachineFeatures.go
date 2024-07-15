package googlesqldatabaseinstance


type GoogleSqlDatabaseInstanceSettingsAdvancedMachineFeatures struct {
	// The number of threads per physical core. Can be 1 or 2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sql_database_instance#threads_per_core GoogleSqlDatabaseInstance#threads_per_core}
	ThreadsPerCore *float64 `field:"optional" json:"threadsPerCore" yaml:"threadsPerCore"`
}

