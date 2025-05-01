package googlemonitoringalertpolicy


type GoogleMonitoringAlertPolicyConditionsConditionSql struct {
	// The Log Analytics SQL query to run, as a string.
	//
	// The query must
	// conform to the required shape. Specifically, the query must not try to
	// filter the input by time.  A filter will automatically be applied
	// to filter the input so that the query receives all rows received
	// since the last time the query was run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_monitoring_alert_policy#query GoogleMonitoringAlertPolicy#query}
	Query *string `field:"required" json:"query" yaml:"query"`
	// boolean_test block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_monitoring_alert_policy#boolean_test GoogleMonitoringAlertPolicy#boolean_test}
	BooleanTest *GoogleMonitoringAlertPolicyConditionsConditionSqlBooleanTest `field:"optional" json:"booleanTest" yaml:"booleanTest"`
	// daily block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_monitoring_alert_policy#daily GoogleMonitoringAlertPolicy#daily}
	Daily *GoogleMonitoringAlertPolicyConditionsConditionSqlDaily `field:"optional" json:"daily" yaml:"daily"`
	// hourly block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_monitoring_alert_policy#hourly GoogleMonitoringAlertPolicy#hourly}
	Hourly *GoogleMonitoringAlertPolicyConditionsConditionSqlHourly `field:"optional" json:"hourly" yaml:"hourly"`
	// minutes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_monitoring_alert_policy#minutes GoogleMonitoringAlertPolicy#minutes}
	Minutes *GoogleMonitoringAlertPolicyConditionsConditionSqlMinutes `field:"optional" json:"minutes" yaml:"minutes"`
	// row_count_test block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_monitoring_alert_policy#row_count_test GoogleMonitoringAlertPolicy#row_count_test}
	RowCountTest *GoogleMonitoringAlertPolicyConditionsConditionSqlRowCountTest `field:"optional" json:"rowCountTest" yaml:"rowCountTest"`
}

