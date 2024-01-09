package slo


type SloObjectiveCountMetricsTotalAzureMonitor struct {
	// Aggregation type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#aggregation Slo#aggregation}
	Aggregation *string `field:"required" json:"aggregation" yaml:"aggregation"`
	// Name of the metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#metric_name Slo#metric_name}
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Name of the added application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#resource_id Slo#resource_id}
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// dimensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#dimensions Slo#dimensions}
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// Namespace of the metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#metric_namespace Slo#metric_namespace}
	MetricNamespace *string `field:"optional" json:"metricNamespace" yaml:"metricNamespace"`
}

