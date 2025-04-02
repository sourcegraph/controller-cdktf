package slo


type SloObjectiveComposite struct {
	// Maximum time for your composite SLO to wait for data from objectives.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#max_delay Slo#max_delay}
	MaxDelay *string `field:"required" json:"maxDelay" yaml:"maxDelay"`
	// components block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#components Slo#components}
	Components *SloObjectiveCompositeComponents `field:"optional" json:"components" yaml:"components"`
}

