package slo


type SloObjectiveCountMetricsTotalGraphite struct {
	// Path to the metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#metric_path Slo#metric_path}
	MetricPath *string `field:"required" json:"metricPath" yaml:"metricPath"`
}

