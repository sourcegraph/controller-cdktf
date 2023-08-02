package googledataformrepository


type GoogleDataformRepositoryWorkspaceCompilationOverrides struct {
	// Optional. The default database (Google Cloud project ID).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.76.0/docs/resources/google_dataform_repository#default_database GoogleDataformRepository#default_database}
	DefaultDatabase *string `field:"optional" json:"defaultDatabase" yaml:"defaultDatabase"`
	// Optional. The suffix that should be appended to all schema (BigQuery dataset ID) names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.76.0/docs/resources/google_dataform_repository#schema_suffix GoogleDataformRepository#schema_suffix}
	SchemaSuffix *string `field:"optional" json:"schemaSuffix" yaml:"schemaSuffix"`
	// Optional. The prefix that should be prepended to all table names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.76.0/docs/resources/google_dataform_repository#table_prefix GoogleDataformRepository#table_prefix}
	TablePrefix *string `field:"optional" json:"tablePrefix" yaml:"tablePrefix"`
}

