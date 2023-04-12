package googlesqldatabaseinstance


type GoogleSqlDatabaseInstanceSettingsInsightsConfig struct {
	// True if Query Insights feature is enabled.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_sql_database_instance#query_insights_enabled GoogleSqlDatabaseInstance#query_insights_enabled}
	QueryInsightsEnabled interface{} `field:"optional" json:"queryInsightsEnabled" yaml:"queryInsightsEnabled"`
	// Maximum query length stored in bytes. Between 256 and 4500. Default to 1024.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_sql_database_instance#query_string_length GoogleSqlDatabaseInstance#query_string_length}
	QueryStringLength *float64 `field:"optional" json:"queryStringLength" yaml:"queryStringLength"`
	// True if Query Insights will record application tags from query when enabled.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_sql_database_instance#record_application_tags GoogleSqlDatabaseInstance#record_application_tags}
	RecordApplicationTags interface{} `field:"optional" json:"recordApplicationTags" yaml:"recordApplicationTags"`
	// True if Query Insights will record client address when enabled.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_sql_database_instance#record_client_address GoogleSqlDatabaseInstance#record_client_address}
	RecordClientAddress interface{} `field:"optional" json:"recordClientAddress" yaml:"recordClientAddress"`
}

