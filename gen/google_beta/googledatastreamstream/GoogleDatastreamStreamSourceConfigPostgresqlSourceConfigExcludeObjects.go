package googledatastreamstream


type GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigExcludeObjects struct {
	// postgresql_schemas block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_datastream_stream#postgresql_schemas GoogleDatastreamStream#postgresql_schemas}
	PostgresqlSchemas interface{} `field:"required" json:"postgresqlSchemas" yaml:"postgresqlSchemas"`
}

