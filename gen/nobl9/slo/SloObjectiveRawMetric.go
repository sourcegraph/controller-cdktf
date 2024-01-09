package slo


type SloObjectiveRawMetric struct {
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#query Slo#query}
	Query interface{} `field:"required" json:"query" yaml:"query"`
}

