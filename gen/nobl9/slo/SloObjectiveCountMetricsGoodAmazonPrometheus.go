package slo


type SloObjectiveCountMetricsGoodAmazonPrometheus struct {
	// Query for the metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#promql Slo#promql}
	Promql *string `field:"required" json:"promql" yaml:"promql"`
}

