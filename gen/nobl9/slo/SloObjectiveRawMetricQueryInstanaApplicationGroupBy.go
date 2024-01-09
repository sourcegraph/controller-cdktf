package slo


type SloObjectiveRawMetricQueryInstanaApplicationGroupBy struct {
	// Group by tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#tag Slo#tag}
	Tag *string `field:"required" json:"tag" yaml:"tag"`
	// Tag entity - one of 'DESTINATION', 'SOURCE', 'NOT_APPLICABLE'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#tag_entity Slo#tag_entity}
	TagEntity *string `field:"required" json:"tagEntity" yaml:"tagEntity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#tag_second_level_key Slo#tag_second_level_key}.
	TagSecondLevelKey *string `field:"optional" json:"tagSecondLevelKey" yaml:"tagSecondLevelKey"`
}

