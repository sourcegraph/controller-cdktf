package spannerinstance


type SpannerInstanceAutoscalingConfigAsymmetricAutoscalingOptionsOverridesAutoscalingLimits struct {
	// The maximum number of nodes for this specific replica.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/spanner_instance#max_nodes SpannerInstance#max_nodes}
	MaxNodes *float64 `field:"optional" json:"maxNodes" yaml:"maxNodes"`
	// The maximum number of processing units for this specific replica.
	//
	// If set, this number should be multiples of 1000 and be greater than or equal to
	// min_processing_units.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/spanner_instance#max_processing_units SpannerInstance#max_processing_units}
	MaxProcessingUnits *float64 `field:"optional" json:"maxProcessingUnits" yaml:"maxProcessingUnits"`
	// The minimum number of nodes for this specific replica.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/spanner_instance#min_nodes SpannerInstance#min_nodes}
	MinNodes *float64 `field:"optional" json:"minNodes" yaml:"minNodes"`
	// The minimum number of processing units for this specific replica. If set, this number should be multiples of 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/spanner_instance#min_processing_units SpannerInstance#min_processing_units}
	MinProcessingUnits *float64 `field:"optional" json:"minProcessingUnits" yaml:"minProcessingUnits"`
}

