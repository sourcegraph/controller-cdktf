package slo


type SloObjectiveCountMetricsTotalGcm struct {
	// Project ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#project_id Slo#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Query for the metrics in PromQL format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#promql Slo#promql}
	Promql *string `field:"optional" json:"promql" yaml:"promql"`
	// Query for the metrics in MQL format ([deprecated](https://cloud.google.com/stackdriver/docs/deprecations/mql)).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#query Slo#query}
	Query *string `field:"optional" json:"query" yaml:"query"`
}

