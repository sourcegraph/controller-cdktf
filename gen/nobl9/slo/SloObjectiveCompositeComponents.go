package slo


type SloObjectiveCompositeComponents struct {
	// objectives block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#objectives Slo#objectives}
	Objectives *SloObjectiveCompositeComponentsObjectives `field:"optional" json:"objectives" yaml:"objectives"`
}

