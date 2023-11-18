package googlespannerinstance


type GoogleSpannerInstanceAutoscalingConfigAutoscalingLimits struct {
	// Specifies maximum number of processing units allocated to the instance.
	//
	// If set, this number should be multiples of 1000 and be greater than or equal to
	// min_processing_units.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.6.0/docs/resources/google_spanner_instance#max_processing_units GoogleSpannerInstance#max_processing_units}
	MaxProcessingUnits *float64 `field:"optional" json:"maxProcessingUnits" yaml:"maxProcessingUnits"`
	// Specifies minimum number of processing units allocated to the instance. If set, this number should be multiples of 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.6.0/docs/resources/google_spanner_instance#min_processing_units GoogleSpannerInstance#min_processing_units}
	MinProcessingUnits *float64 `field:"optional" json:"minProcessingUnits" yaml:"minProcessingUnits"`
}

