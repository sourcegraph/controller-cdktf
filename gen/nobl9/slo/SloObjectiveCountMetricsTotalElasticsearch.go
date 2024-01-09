package slo


type SloObjectiveCountMetricsTotalElasticsearch struct {
	// Index of metrics we want to query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#index Slo#index}
	Index *string `field:"required" json:"index" yaml:"index"`
	// Query for the metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#query Slo#query}
	Query *string `field:"required" json:"query" yaml:"query"`
}

