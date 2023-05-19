package googledatastreamstream


type GoogleDatastreamStreamBackfillAll struct {
	// mysql_excluded_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_datastream_stream#mysql_excluded_objects GoogleDatastreamStream#mysql_excluded_objects}
	MysqlExcludedObjects *GoogleDatastreamStreamBackfillAllMysqlExcludedObjects `field:"optional" json:"mysqlExcludedObjects" yaml:"mysqlExcludedObjects"`
	// oracle_excluded_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_datastream_stream#oracle_excluded_objects GoogleDatastreamStream#oracle_excluded_objects}
	OracleExcludedObjects *GoogleDatastreamStreamBackfillAllOracleExcludedObjects `field:"optional" json:"oracleExcludedObjects" yaml:"oracleExcludedObjects"`
	// postgresql_excluded_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_datastream_stream#postgresql_excluded_objects GoogleDatastreamStream#postgresql_excluded_objects}
	PostgresqlExcludedObjects *GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjects `field:"optional" json:"postgresqlExcludedObjects" yaml:"postgresqlExcludedObjects"`
}

