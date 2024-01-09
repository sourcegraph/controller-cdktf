package slo


type SloObjectiveCountMetricsBadSumologic struct {
	// Query for the metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#query Slo#query}
	Query *string `field:"required" json:"query" yaml:"query"`
	// Sumologic source - metrics or logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#type Slo#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Period of data aggregation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#quantization Slo#quantization}
	Quantization *string `field:"optional" json:"quantization" yaml:"quantization"`
	// Aggregation function - avg, sum, min, max, count, none.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#rollup Slo#rollup}
	Rollup *string `field:"optional" json:"rollup" yaml:"rollup"`
}

