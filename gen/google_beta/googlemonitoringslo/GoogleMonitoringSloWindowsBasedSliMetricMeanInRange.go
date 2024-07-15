package googlemonitoringslo


type GoogleMonitoringSloWindowsBasedSliMetricMeanInRange struct {
	// range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_slo#range GoogleMonitoringSlo#range}
	Range *GoogleMonitoringSloWindowsBasedSliMetricMeanInRangeRange `field:"required" json:"range" yaml:"range"`
	// A [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters) specifying the TimeSeries to use for evaluating window The provided TimeSeries must have ValueType = INT64 or ValueType = DOUBLE and MetricKind = GAUGE. Mean value 'X' should satisfy 'range.min <= X <= range.max' under good service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_slo#time_series GoogleMonitoringSlo#time_series}
	TimeSeries *string `field:"required" json:"timeSeries" yaml:"timeSeries"`
}

