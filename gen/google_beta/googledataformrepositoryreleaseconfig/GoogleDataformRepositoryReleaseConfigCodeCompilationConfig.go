package googledataformrepositoryreleaseconfig


type GoogleDataformRepositoryReleaseConfigCodeCompilationConfig struct {
	// Optional. The default schema (BigQuery dataset ID) for assertions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository_release_config#assertion_schema GoogleDataformRepositoryReleaseConfig#assertion_schema}
	AssertionSchema *string `field:"optional" json:"assertionSchema" yaml:"assertionSchema"`
	// Optional. The suffix that should be appended to all database (Google Cloud project ID) names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository_release_config#database_suffix GoogleDataformRepositoryReleaseConfig#database_suffix}
	DatabaseSuffix *string `field:"optional" json:"databaseSuffix" yaml:"databaseSuffix"`
	// Optional. The default database (Google Cloud project ID).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository_release_config#default_database GoogleDataformRepositoryReleaseConfig#default_database}
	DefaultDatabase *string `field:"optional" json:"defaultDatabase" yaml:"defaultDatabase"`
	// Optional.
	//
	// The default BigQuery location to use. Defaults to "US".
	// See the BigQuery docs for a full list of locations: https://cloud.google.com/bigquery/docs/locations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository_release_config#default_location GoogleDataformRepositoryReleaseConfig#default_location}
	DefaultLocation *string `field:"optional" json:"defaultLocation" yaml:"defaultLocation"`
	// Optional. The default schema (BigQuery dataset ID).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository_release_config#default_schema GoogleDataformRepositoryReleaseConfig#default_schema}
	DefaultSchema *string `field:"optional" json:"defaultSchema" yaml:"defaultSchema"`
	// Optional. The suffix that should be appended to all schema (BigQuery dataset ID) names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository_release_config#schema_suffix GoogleDataformRepositoryReleaseConfig#schema_suffix}
	SchemaSuffix *string `field:"optional" json:"schemaSuffix" yaml:"schemaSuffix"`
	// Optional. The prefix that should be prepended to all table names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository_release_config#table_prefix GoogleDataformRepositoryReleaseConfig#table_prefix}
	TablePrefix *string `field:"optional" json:"tablePrefix" yaml:"tablePrefix"`
	// Optional.
	//
	// User-defined variables that are made available to project code during compilation.
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository_release_config#vars GoogleDataformRepositoryReleaseConfig#vars}
	Vars *map[string]*string `field:"optional" json:"vars" yaml:"vars"`
}

