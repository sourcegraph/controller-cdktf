package slo


type SloObjectiveRawMetricQueryGcm struct {
	// Project ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#project_id Slo#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Query for the metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#query Slo#query}
	Query *string `field:"required" json:"query" yaml:"query"`
}

