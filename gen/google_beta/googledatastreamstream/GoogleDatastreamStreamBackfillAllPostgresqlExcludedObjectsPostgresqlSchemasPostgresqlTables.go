package googledatastreamstream


type GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTables struct {
	// Table name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_stream#table GoogleDatastreamStream#table}
	Table *string `field:"required" json:"table" yaml:"table"`
	// postgresql_columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_stream#postgresql_columns GoogleDatastreamStream#postgresql_columns}
	PostgresqlColumns interface{} `field:"optional" json:"postgresqlColumns" yaml:"postgresqlColumns"`
}

