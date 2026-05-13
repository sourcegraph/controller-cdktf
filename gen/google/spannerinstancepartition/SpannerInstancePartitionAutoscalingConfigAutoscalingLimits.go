package spannerinstancepartition


type SpannerInstancePartitionAutoscalingConfigAutoscalingLimits struct {
	// Specifies maximum number of nodes allocated to the instance partition.
	//
	// If set, this number
	// should be greater than or equal to min_nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/spanner_instance_partition#max_nodes SpannerInstancePartition#max_nodes}
	MaxNodes *float64 `field:"optional" json:"maxNodes" yaml:"maxNodes"`
	// Specifies maximum number of processing units allocated to the instance partition.
	//
	// If set, this number should be multiples of 1000 and be greater than or equal to
	// min_processing_units.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/spanner_instance_partition#max_processing_units SpannerInstancePartition#max_processing_units}
	MaxProcessingUnits *float64 `field:"optional" json:"maxProcessingUnits" yaml:"maxProcessingUnits"`
	// Specifies number of nodes allocated to the instance partition.
	//
	// If set, this number
	// should be greater than or equal to 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/spanner_instance_partition#min_nodes SpannerInstancePartition#min_nodes}
	MinNodes *float64 `field:"optional" json:"minNodes" yaml:"minNodes"`
	// Specifies minimum number of processing units allocated to the instance partition. If set, this number should be multiples of 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/spanner_instance_partition#min_processing_units SpannerInstancePartition#min_processing_units}
	MinProcessingUnits *float64 `field:"optional" json:"minProcessingUnits" yaml:"minProcessingUnits"`
}

