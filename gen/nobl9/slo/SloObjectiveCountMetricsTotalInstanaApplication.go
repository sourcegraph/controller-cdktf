package slo


type SloObjectiveCountMetricsTotalInstanaApplication struct {
	// Depends on the value specified for 'metric_id'- more info in N9 docs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#aggregation Slo#aggregation}
	Aggregation *string `field:"required" json:"aggregation" yaml:"aggregation"`
	// API query user passes in a JSON format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#api_query Slo#api_query}
	ApiQuery *string `field:"required" json:"apiQuery" yaml:"apiQuery"`
	// group_by block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#group_by Slo#group_by}
	GroupBy interface{} `field:"required" json:"groupBy" yaml:"groupBy"`
	// Metric ID one of 'calls', 'erroneousCalls', 'errors', 'latency'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#metric_id Slo#metric_id}
	MetricId *string `field:"required" json:"metricId" yaml:"metricId"`
	// Include internal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#include_internal Slo#include_internal}
	IncludeInternal interface{} `field:"optional" json:"includeInternal" yaml:"includeInternal"`
	// Include synthetic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#include_synthetic Slo#include_synthetic}
	IncludeSynthetic interface{} `field:"optional" json:"includeSynthetic" yaml:"includeSynthetic"`
}

