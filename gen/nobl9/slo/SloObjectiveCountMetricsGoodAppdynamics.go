package slo


type SloObjectiveCountMetricsGoodAppdynamics struct {
	// Name of the added application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#application_name Slo#application_name}
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
	// Path to the metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#metric_path Slo#metric_path}
	MetricPath *string `field:"required" json:"metricPath" yaml:"metricPath"`
}

