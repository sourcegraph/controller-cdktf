package googledatastreamstream


type GoogleDatastreamStreamSourceConfig struct {
	// Source connection profile resource. Format: projects/{project}/locations/{location}/connectionProfiles/{name}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_stream#source_connection_profile GoogleDatastreamStream#source_connection_profile}
	SourceConnectionProfile *string `field:"required" json:"sourceConnectionProfile" yaml:"sourceConnectionProfile"`
	// mysql_source_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_stream#mysql_source_config GoogleDatastreamStream#mysql_source_config}
	MysqlSourceConfig *GoogleDatastreamStreamSourceConfigMysqlSourceConfig `field:"optional" json:"mysqlSourceConfig" yaml:"mysqlSourceConfig"`
	// oracle_source_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_stream#oracle_source_config GoogleDatastreamStream#oracle_source_config}
	OracleSourceConfig *GoogleDatastreamStreamSourceConfigOracleSourceConfig `field:"optional" json:"oracleSourceConfig" yaml:"oracleSourceConfig"`
	// postgresql_source_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_stream#postgresql_source_config GoogleDatastreamStream#postgresql_source_config}
	PostgresqlSourceConfig *GoogleDatastreamStreamSourceConfigPostgresqlSourceConfig `field:"optional" json:"postgresqlSourceConfig" yaml:"postgresqlSourceConfig"`
	// sql_server_source_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_stream#sql_server_source_config GoogleDatastreamStream#sql_server_source_config}
	SqlServerSourceConfig *GoogleDatastreamStreamSourceConfigSqlServerSourceConfig `field:"optional" json:"sqlServerSourceConfig" yaml:"sqlServerSourceConfig"`
}

