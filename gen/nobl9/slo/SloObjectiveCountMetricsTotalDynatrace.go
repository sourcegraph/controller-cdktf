package slo


type SloObjectiveCountMetricsTotalDynatrace struct {
	// Selector for the metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#metric_selector Slo#metric_selector}
	MetricSelector *string `field:"required" json:"metricSelector" yaml:"metricSelector"`
}

