package slo


type SloObjectiveRawMetricQueryInstana struct {
	// Instana metric type 'application' or 'infrastructure'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#metric_type Slo#metric_type}
	MetricType *string `field:"required" json:"metricType" yaml:"metricType"`
	// application block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#application Slo#application}
	Application interface{} `field:"optional" json:"application" yaml:"application"`
	// infrastructure block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#infrastructure Slo#infrastructure}
	Infrastructure interface{} `field:"optional" json:"infrastructure" yaml:"infrastructure"`
}

