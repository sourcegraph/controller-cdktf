package slo


type SloObjectiveCountMetricsBadBigquery struct {
	// Location of you BigQuery.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#location Slo#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Project ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#project_id Slo#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Query for the metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#query Slo#query}
	Query *string `field:"required" json:"query" yaml:"query"`
}

