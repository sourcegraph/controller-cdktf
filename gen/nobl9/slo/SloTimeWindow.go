package slo


type SloTimeWindow struct {
	// Count of the time unit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#count Slo#count}
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Unit of time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#unit Slo#unit}
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// calendar block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#calendar Slo#calendar}
	Calendar interface{} `field:"optional" json:"calendar" yaml:"calendar"`
	// Is the window moving or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#is_rolling Slo#is_rolling}
	IsRolling interface{} `field:"optional" json:"isRolling" yaml:"isRolling"`
}

