package googlemonitoringalertpolicy


type GoogleMonitoringAlertPolicyConditionsConditionSqlMinutes struct {
	// Number of minutes between runs.
	//
	// The interval must be greater than or
	// equal to 5 minutes and less than or equal to 1440 minutes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_monitoring_alert_policy#periodicity GoogleMonitoringAlertPolicy#periodicity}
	Periodicity *float64 `field:"required" json:"periodicity" yaml:"periodicity"`
}

