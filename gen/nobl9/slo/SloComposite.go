package slo


type SloComposite struct {
	// The numeric target for your objective.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#target Slo#target}
	Target *float64 `field:"required" json:"target" yaml:"target"`
	// burn_rate_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#burn_rate_condition Slo#burn_rate_condition}
	BurnRateCondition interface{} `field:"optional" json:"burnRateCondition" yaml:"burnRateCondition"`
}

