package datastreamstream

type DatastreamStreamBackfillAllPostgresqlExcludedObjects struct {
	// postgresql_schemas block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/datastream_stream#postgresql_schemas DatastreamStream#postgresql_schemas}
	PostgresqlSchemas any `field:"required" json:"postgresqlSchemas" yaml:"postgresqlSchemas"`
}
