package googlemonitoringalertpolicy


type GoogleMonitoringAlertPolicyConditionsConditionSqlDaily struct {
	// The number of days between runs.
	//
	// Must be greater than or equal
	// to 1 day and less than or equal to 30 days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_monitoring_alert_policy#periodicity GoogleMonitoringAlertPolicy#periodicity}
	Periodicity *float64 `field:"required" json:"periodicity" yaml:"periodicity"`
	// execution_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_monitoring_alert_policy#execution_time GoogleMonitoringAlertPolicy#execution_time}
	ExecutionTime *GoogleMonitoringAlertPolicyConditionsConditionSqlDailyExecutionTime `field:"optional" json:"executionTime" yaml:"executionTime"`
}

