package slo


type SloObjectiveRawMetricQueryHoneycomb struct {
	// Calculation type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#calculation Slo#calculation}
	Calculation *string `field:"required" json:"calculation" yaml:"calculation"`
	// Column name - required for all calculation types besides 'CONCURRENCY' and 'COUNT'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#attribute Slo#attribute}
	Attribute *string `field:"optional" json:"attribute" yaml:"attribute"`
}

