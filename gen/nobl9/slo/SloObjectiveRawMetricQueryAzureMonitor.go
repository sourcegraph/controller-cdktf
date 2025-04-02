package slo


type SloObjectiveRawMetricQueryAzureMonitor struct {
	// Specifies source: 'metrics' or 'logs'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#data_type Slo#data_type}
	DataType *string `field:"required" json:"dataType" yaml:"dataType"`
	// Aggregation type [Required for metrics].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#aggregation Slo#aggregation}
	Aggregation *string `field:"optional" json:"aggregation" yaml:"aggregation"`
	// dimensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#dimensions Slo#dimensions}
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// Logs query in Kusto Query Language [Required for logs].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#kql_query Slo#kql_query}
	KqlQuery *string `field:"optional" json:"kqlQuery" yaml:"kqlQuery"`
	// Name of the metric [Required for metrics].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#metric_name Slo#metric_name}
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// Namespace of the metric [Optional for metrics].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#metric_namespace Slo#metric_namespace}
	MetricNamespace *string `field:"optional" json:"metricNamespace" yaml:"metricNamespace"`
	// Identifier of the Azure Cloud resource [Required for metrics].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#resource_id Slo#resource_id}
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
	// workspace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#workspace Slo#workspace}
	Workspace interface{} `field:"optional" json:"workspace" yaml:"workspace"`
}

