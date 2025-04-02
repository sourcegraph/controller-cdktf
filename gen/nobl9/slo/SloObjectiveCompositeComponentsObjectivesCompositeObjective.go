package slo


type SloObjectiveCompositeComponentsObjectivesCompositeObjective struct {
	// SLO objective name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#objective Slo#objective}
	Objective *string `field:"required" json:"objective" yaml:"objective"`
	// Project name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#project Slo#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// SLO name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#slo Slo#slo}
	Slo *string `field:"required" json:"slo" yaml:"slo"`
	// Weights determine each component’s contribution to the composite SLO.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#weight Slo#weight}
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
	// Defines how to treat missing component data on `max_delay` expiry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#when_delayed Slo#when_delayed}
	WhenDelayed *string `field:"required" json:"whenDelayed" yaml:"whenDelayed"`
}

