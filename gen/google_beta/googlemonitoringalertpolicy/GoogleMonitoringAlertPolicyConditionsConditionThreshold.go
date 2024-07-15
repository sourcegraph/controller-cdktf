package googlemonitoringalertpolicy


type GoogleMonitoringAlertPolicyConditionsConditionThreshold struct {
	// The comparison to apply between the time series (indicated by filter and aggregation) and the threshold (indicated by threshold_value).
	//
	// The comparison is applied
	// on each time series, with the time series on
	// the left-hand side and the threshold on the
	// right-hand side. Only COMPARISON_LT and
	// COMPARISON_GT are supported currently. Possible values: ["COMPARISON_GT", "COMPARISON_GE", "COMPARISON_LT", "COMPARISON_LE", "COMPARISON_EQ", "COMPARISON_NE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_alert_policy#comparison GoogleMonitoringAlertPolicy#comparison}
	Comparison *string `field:"required" json:"comparison" yaml:"comparison"`
	// The amount of time that a time series must violate the threshold to be considered failing.
	//
	// Currently, only values that are a
	// multiple of a minute--e.g., 0, 60, 120, or
	// 300 seconds--are supported. If an invalid
	// value is given, an error will be returned.
	// When choosing a duration, it is useful to
	// keep in mind the frequency of the underlying
	// time series data (which may also be affected
	// by any alignments specified in the
	// aggregations field); a good duration is long
	// enough so that a single outlier does not
	// generate spurious alerts, but short enough
	// that unhealthy states are detected and
	// alerted on quickly.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_alert_policy#duration GoogleMonitoringAlertPolicy#duration}
	Duration *string `field:"required" json:"duration" yaml:"duration"`
	// aggregations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_alert_policy#aggregations GoogleMonitoringAlertPolicy#aggregations}
	Aggregations interface{} `field:"optional" json:"aggregations" yaml:"aggregations"`
	// denominator_aggregations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_alert_policy#denominator_aggregations GoogleMonitoringAlertPolicy#denominator_aggregations}
	DenominatorAggregations interface{} `field:"optional" json:"denominatorAggregations" yaml:"denominatorAggregations"`
	// A filter that identifies a time series that should be used as the denominator of a ratio that will be compared with the threshold.
	//
	// If
	// a denominator_filter is specified, the time
	// series specified by the filter field will be
	// used as the numerator.The filter is similar
	// to the one that is specified in the
	// MetricService.ListTimeSeries request (that
	// call is useful to verify the time series
	// that will be retrieved / processed) and must
	// specify the metric type and optionally may
	// contain restrictions on resource type,
	// resource labels, and metric labels. This
	// field may not exceed 2048 Unicode characters
	// in length.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_alert_policy#denominator_filter GoogleMonitoringAlertPolicy#denominator_filter}
	DenominatorFilter *string `field:"optional" json:"denominatorFilter" yaml:"denominatorFilter"`
	// A condition control that determines how metric-threshold conditions are evaluated when data stops arriving. Possible values: ["EVALUATION_MISSING_DATA_INACTIVE", "EVALUATION_MISSING_DATA_ACTIVE", "EVALUATION_MISSING_DATA_NO_OP"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_alert_policy#evaluation_missing_data GoogleMonitoringAlertPolicy#evaluation_missing_data}
	EvaluationMissingData *string `field:"optional" json:"evaluationMissingData" yaml:"evaluationMissingData"`
	// A filter that identifies which time series should be compared with the threshold.The filter is similar to the one that is specified in the MetricService.ListTimeSeries request (that call is useful to verify the time series that will be retrieved / processed) and must specify the metric type and optionally may contain restrictions on resource type, resource labels, and metric labels. This field may not exceed 2048 Unicode characters in length.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_alert_policy#filter GoogleMonitoringAlertPolicy#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// forecast_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_alert_policy#forecast_options GoogleMonitoringAlertPolicy#forecast_options}
	ForecastOptions *GoogleMonitoringAlertPolicyConditionsConditionThresholdForecastOptions `field:"optional" json:"forecastOptions" yaml:"forecastOptions"`
	// A value against which to compare the time series.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_alert_policy#threshold_value GoogleMonitoringAlertPolicy#threshold_value}
	ThresholdValue *float64 `field:"optional" json:"thresholdValue" yaml:"thresholdValue"`
	// trigger block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_alert_policy#trigger GoogleMonitoringAlertPolicy#trigger}
	Trigger *GoogleMonitoringAlertPolicyConditionsConditionThresholdTrigger `field:"optional" json:"trigger" yaml:"trigger"`
}

