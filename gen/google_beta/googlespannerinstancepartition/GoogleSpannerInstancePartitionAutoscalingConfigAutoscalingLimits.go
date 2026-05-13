package googlespannerinstancepartition


type GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimits struct {
	// Specifies maximum number of nodes allocated to the instance partition.
	//
	// If set, this number
	// should be greater than or equal to min_nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_spanner_instance_partition#max_nodes GoogleSpannerInstancePartition#max_nodes}
	MaxNodes *float64 `field:"optional" json:"maxNodes" yaml:"maxNodes"`
	// Specifies maximum number of processing units allocated to the instance partition.
	//
	// If set, this number should be multiples of 1000 and be greater than or equal to
	// min_processing_units.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_spanner_instance_partition#max_processing_units GoogleSpannerInstancePartition#max_processing_units}
	MaxProcessingUnits *float64 `field:"optional" json:"maxProcessingUnits" yaml:"maxProcessingUnits"`
	// Specifies number of nodes allocated to the instance partition.
	//
	// If set, this number
	// should be greater than or equal to 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_spanner_instance_partition#min_nodes GoogleSpannerInstancePartition#min_nodes}
	MinNodes *float64 `field:"optional" json:"minNodes" yaml:"minNodes"`
	// Specifies minimum number of processing units allocated to the instance partition. If set, this number should be multiples of 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_spanner_instance_partition#min_processing_units GoogleSpannerInstancePartition#min_processing_units}
	MinProcessingUnits *float64 `field:"optional" json:"minProcessingUnits" yaml:"minProcessingUnits"`
}

