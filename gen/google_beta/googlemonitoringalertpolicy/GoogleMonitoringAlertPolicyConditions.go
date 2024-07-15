package googlemonitoringalertpolicy


type GoogleMonitoringAlertPolicyConditions struct {
	// A short name or phrase used to identify the condition in dashboards, notifications, and incidents.
	//
	// To avoid confusion, don't use the same
	// display name for multiple conditions in the same
	// policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_alert_policy#display_name GoogleMonitoringAlertPolicy#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// condition_absent block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_alert_policy#condition_absent GoogleMonitoringAlertPolicy#condition_absent}
	ConditionAbsent *GoogleMonitoringAlertPolicyConditionsConditionAbsent `field:"optional" json:"conditionAbsent" yaml:"conditionAbsent"`
	// condition_matched_log block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_alert_policy#condition_matched_log GoogleMonitoringAlertPolicy#condition_matched_log}
	ConditionMatchedLog *GoogleMonitoringAlertPolicyConditionsConditionMatchedLog `field:"optional" json:"conditionMatchedLog" yaml:"conditionMatchedLog"`
	// condition_monitoring_query_language block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_alert_policy#condition_monitoring_query_language GoogleMonitoringAlertPolicy#condition_monitoring_query_language}
	ConditionMonitoringQueryLanguage *GoogleMonitoringAlertPolicyConditionsConditionMonitoringQueryLanguage `field:"optional" json:"conditionMonitoringQueryLanguage" yaml:"conditionMonitoringQueryLanguage"`
	// condition_prometheus_query_language block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_alert_policy#condition_prometheus_query_language GoogleMonitoringAlertPolicy#condition_prometheus_query_language}
	ConditionPrometheusQueryLanguage *GoogleMonitoringAlertPolicyConditionsConditionPrometheusQueryLanguage `field:"optional" json:"conditionPrometheusQueryLanguage" yaml:"conditionPrometheusQueryLanguage"`
	// condition_threshold block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_alert_policy#condition_threshold GoogleMonitoringAlertPolicy#condition_threshold}
	ConditionThreshold *GoogleMonitoringAlertPolicyConditionsConditionThreshold `field:"optional" json:"conditionThreshold" yaml:"conditionThreshold"`
}

