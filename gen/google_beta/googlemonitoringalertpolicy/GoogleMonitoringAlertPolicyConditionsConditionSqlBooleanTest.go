package googlemonitoringalertpolicy


type GoogleMonitoringAlertPolicyConditionsConditionSqlBooleanTest struct {
	// The name of the column containing the boolean value.
	//
	// If the value in a row is
	// NULL, that row is ignored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_monitoring_alert_policy#column GoogleMonitoringAlertPolicy#column}
	Column *string `field:"required" json:"column" yaml:"column"`
}

