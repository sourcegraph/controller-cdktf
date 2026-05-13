package datastreamstream


type DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifier struct {
	// mongodb_identifier block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#mongodb_identifier DatastreamStream#mongodb_identifier}
	MongodbIdentifier *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierMongodbIdentifier `field:"optional" json:"mongodbIdentifier" yaml:"mongodbIdentifier"`
	// mysql_identifier block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#mysql_identifier DatastreamStream#mysql_identifier}
	MysqlIdentifier *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierMysqlIdentifier `field:"optional" json:"mysqlIdentifier" yaml:"mysqlIdentifier"`
	// oracle_identifier block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#oracle_identifier DatastreamStream#oracle_identifier}
	OracleIdentifier *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOracleIdentifier `field:"optional" json:"oracleIdentifier" yaml:"oracleIdentifier"`
	// postgresql_identifier block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#postgresql_identifier DatastreamStream#postgresql_identifier}
	PostgresqlIdentifier *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierPostgresqlIdentifier `field:"optional" json:"postgresqlIdentifier" yaml:"postgresqlIdentifier"`
	// salesforce_identifier block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#salesforce_identifier DatastreamStream#salesforce_identifier}
	SalesforceIdentifier *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSalesforceIdentifier `field:"optional" json:"salesforceIdentifier" yaml:"salesforceIdentifier"`
	// spanner_identifier block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#spanner_identifier DatastreamStream#spanner_identifier}
	SpannerIdentifier *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSpannerIdentifier `field:"optional" json:"spannerIdentifier" yaml:"spannerIdentifier"`
	// sql_server_identifier block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#sql_server_identifier DatastreamStream#sql_server_identifier}
	SqlServerIdentifier *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSqlServerIdentifier `field:"optional" json:"sqlServerIdentifier" yaml:"sqlServerIdentifier"`
}

