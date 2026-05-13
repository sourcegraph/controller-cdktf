package googledatastreamstream


type GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifier struct {
	// mongodb_identifier block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_stream#mongodb_identifier GoogleDatastreamStream#mongodb_identifier}
	MongodbIdentifier *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierMongodbIdentifier `field:"optional" json:"mongodbIdentifier" yaml:"mongodbIdentifier"`
	// mysql_identifier block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_stream#mysql_identifier GoogleDatastreamStream#mysql_identifier}
	MysqlIdentifier *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierMysqlIdentifier `field:"optional" json:"mysqlIdentifier" yaml:"mysqlIdentifier"`
	// oracle_identifier block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_stream#oracle_identifier GoogleDatastreamStream#oracle_identifier}
	OracleIdentifier *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOracleIdentifier `field:"optional" json:"oracleIdentifier" yaml:"oracleIdentifier"`
	// postgresql_identifier block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_stream#postgresql_identifier GoogleDatastreamStream#postgresql_identifier}
	PostgresqlIdentifier *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierPostgresqlIdentifier `field:"optional" json:"postgresqlIdentifier" yaml:"postgresqlIdentifier"`
	// salesforce_identifier block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_stream#salesforce_identifier GoogleDatastreamStream#salesforce_identifier}
	SalesforceIdentifier *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSalesforceIdentifier `field:"optional" json:"salesforceIdentifier" yaml:"salesforceIdentifier"`
	// spanner_identifier block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_stream#spanner_identifier GoogleDatastreamStream#spanner_identifier}
	SpannerIdentifier *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSpannerIdentifier `field:"optional" json:"spannerIdentifier" yaml:"spannerIdentifier"`
	// sql_server_identifier block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_stream#sql_server_identifier GoogleDatastreamStream#sql_server_identifier}
	SqlServerIdentifier *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSqlServerIdentifier `field:"optional" json:"sqlServerIdentifier" yaml:"sqlServerIdentifier"`
}

