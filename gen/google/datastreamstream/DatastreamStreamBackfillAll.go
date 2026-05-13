package datastreamstream


type DatastreamStreamBackfillAll struct {
	// mongodb_excluded_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#mongodb_excluded_objects DatastreamStream#mongodb_excluded_objects}
	MongodbExcludedObjects *DatastreamStreamBackfillAllMongodbExcludedObjects `field:"optional" json:"mongodbExcludedObjects" yaml:"mongodbExcludedObjects"`
	// mysql_excluded_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#mysql_excluded_objects DatastreamStream#mysql_excluded_objects}
	MysqlExcludedObjects *DatastreamStreamBackfillAllMysqlExcludedObjects `field:"optional" json:"mysqlExcludedObjects" yaml:"mysqlExcludedObjects"`
	// oracle_excluded_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#oracle_excluded_objects DatastreamStream#oracle_excluded_objects}
	OracleExcludedObjects *DatastreamStreamBackfillAllOracleExcludedObjects `field:"optional" json:"oracleExcludedObjects" yaml:"oracleExcludedObjects"`
	// postgresql_excluded_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#postgresql_excluded_objects DatastreamStream#postgresql_excluded_objects}
	PostgresqlExcludedObjects *DatastreamStreamBackfillAllPostgresqlExcludedObjects `field:"optional" json:"postgresqlExcludedObjects" yaml:"postgresqlExcludedObjects"`
	// salesforce_excluded_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#salesforce_excluded_objects DatastreamStream#salesforce_excluded_objects}
	SalesforceExcludedObjects *DatastreamStreamBackfillAllSalesforceExcludedObjects `field:"optional" json:"salesforceExcludedObjects" yaml:"salesforceExcludedObjects"`
	// spanner_excluded_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#spanner_excluded_objects DatastreamStream#spanner_excluded_objects}
	SpannerExcludedObjects *DatastreamStreamBackfillAllSpannerExcludedObjects `field:"optional" json:"spannerExcludedObjects" yaml:"spannerExcludedObjects"`
	// sql_server_excluded_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#sql_server_excluded_objects DatastreamStream#sql_server_excluded_objects}
	SqlServerExcludedObjects *DatastreamStreamBackfillAllSqlServerExcludedObjects `field:"optional" json:"sqlServerExcludedObjects" yaml:"sqlServerExcludedObjects"`
}

