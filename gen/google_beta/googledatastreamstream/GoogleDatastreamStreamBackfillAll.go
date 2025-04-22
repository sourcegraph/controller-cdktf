package googledatastreamstream


type GoogleDatastreamStreamBackfillAll struct {
	// mysql_excluded_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_datastream_stream#mysql_excluded_objects GoogleDatastreamStream#mysql_excluded_objects}
	MysqlExcludedObjects *GoogleDatastreamStreamBackfillAllMysqlExcludedObjects `field:"optional" json:"mysqlExcludedObjects" yaml:"mysqlExcludedObjects"`
	// oracle_excluded_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_datastream_stream#oracle_excluded_objects GoogleDatastreamStream#oracle_excluded_objects}
	OracleExcludedObjects *GoogleDatastreamStreamBackfillAllOracleExcludedObjects `field:"optional" json:"oracleExcludedObjects" yaml:"oracleExcludedObjects"`
	// postgresql_excluded_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_datastream_stream#postgresql_excluded_objects GoogleDatastreamStream#postgresql_excluded_objects}
	PostgresqlExcludedObjects *GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjects `field:"optional" json:"postgresqlExcludedObjects" yaml:"postgresqlExcludedObjects"`
	// salesforce_excluded_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_datastream_stream#salesforce_excluded_objects GoogleDatastreamStream#salesforce_excluded_objects}
	SalesforceExcludedObjects *GoogleDatastreamStreamBackfillAllSalesforceExcludedObjects `field:"optional" json:"salesforceExcludedObjects" yaml:"salesforceExcludedObjects"`
	// sql_server_excluded_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_datastream_stream#sql_server_excluded_objects GoogleDatastreamStream#sql_server_excluded_objects}
	SqlServerExcludedObjects *GoogleDatastreamStreamBackfillAllSqlServerExcludedObjects `field:"optional" json:"sqlServerExcludedObjects" yaml:"sqlServerExcludedObjects"`
}

