package slo


type SloObjectiveRawMetricQueryDatadog struct {
	// Query for the metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#query Slo#query}
	Query *string `field:"required" json:"query" yaml:"query"`
}

