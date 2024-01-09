package slo


type SloCompositeBurnRateCondition struct {
	// Type of logical operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#op Slo#op}
	Op *string `field:"required" json:"op" yaml:"op"`
	// Burn rate value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#value Slo#value}
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

