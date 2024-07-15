package googlemonitoringalertpolicy


type GoogleMonitoringAlertPolicyConditionsConditionAbsent struct {
	// The amount of time that a time series must fail to report new data to be considered failing.
	//
	// Currently, only values that are a
	// multiple of a minute--e.g. 60s, 120s, or 300s
	// --are supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_alert_policy#duration GoogleMonitoringAlertPolicy#duration}
	Duration *string `field:"required" json:"duration" yaml:"duration"`
	// aggregations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_alert_policy#aggregations GoogleMonitoringAlertPolicy#aggregations}
	Aggregations interface{} `field:"optional" json:"aggregations" yaml:"aggregations"`
	// A filter that identifies which time series should be compared with the threshold.The filter is similar to the one that is specified in the MetricService.ListTimeSeries request (that call is useful to verify the time series that will be retrieved / processed) and must specify the metric type and optionally may contain restrictions on resource type, resource labels, and metric labels. This field may not exceed 2048 Unicode characters in length.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_alert_policy#filter GoogleMonitoringAlertPolicy#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// trigger block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_alert_policy#trigger GoogleMonitoringAlertPolicy#trigger}
	Trigger *GoogleMonitoringAlertPolicyConditionsConditionAbsentTrigger `field:"optional" json:"trigger" yaml:"trigger"`
}

