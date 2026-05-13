package googlesqldatabaseinstance


type GoogleSqlDatabaseInstanceSettingsReadPoolAutoScaleConfigTargetMetrics struct {
	// Metric name for Read Pool Auto Scale.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_sql_database_instance#metric GoogleSqlDatabaseInstance#metric}
	Metric *string `field:"optional" json:"metric" yaml:"metric"`
	// Target value for Read Pool Auto Scale.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_sql_database_instance#target_value GoogleSqlDatabaseInstance#target_value}
	TargetValue *float64 `field:"optional" json:"targetValue" yaml:"targetValue"`
}

