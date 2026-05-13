package sqldatabaseinstance


type SqlDatabaseInstanceSettingsReadPoolAutoScaleConfigTargetMetrics struct {
	// Metric name for Read Pool Auto Scale.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/sql_database_instance#metric SqlDatabaseInstance#metric}
	Metric *string `field:"optional" json:"metric" yaml:"metric"`
	// Target value for Read Pool Auto Scale.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/sql_database_instance#target_value SqlDatabaseInstance#target_value}
	TargetValue *float64 `field:"optional" json:"targetValue" yaml:"targetValue"`
}

