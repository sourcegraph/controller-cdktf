package slo


type SloObjectiveCountMetricsTotalCloudwatchDimensions struct {
	// Name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#name Slo#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#value Slo#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

