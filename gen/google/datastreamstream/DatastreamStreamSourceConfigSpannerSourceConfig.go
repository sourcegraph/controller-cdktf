package datastreamstream


type DatastreamStreamSourceConfigSpannerSourceConfig struct {
	// Whether to use DataBoost for backfill queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#backfill_data_boost_enabled DatastreamStream#backfill_data_boost_enabled}
	BackfillDataBoostEnabled interface{} `field:"optional" json:"backfillDataBoostEnabled" yaml:"backfillDataBoostEnabled"`
	// The Spanner change stream name to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#change_stream_name DatastreamStream#change_stream_name}
	ChangeStreamName *string `field:"optional" json:"changeStreamName" yaml:"changeStreamName"`
	// exclude_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#exclude_objects DatastreamStream#exclude_objects}
	ExcludeObjects *DatastreamStreamSourceConfigSpannerSourceConfigExcludeObjects `field:"optional" json:"excludeObjects" yaml:"excludeObjects"`
	// The FGAC role to use for Spanner queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#fgac_role DatastreamStream#fgac_role}
	FgacRole *string `field:"optional" json:"fgacRole" yaml:"fgacRole"`
	// include_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#include_objects DatastreamStream#include_objects}
	IncludeObjects *DatastreamStreamSourceConfigSpannerSourceConfigIncludeObjects `field:"optional" json:"includeObjects" yaml:"includeObjects"`
	// Max concurrent backfill tasks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#max_concurrent_backfill_tasks DatastreamStream#max_concurrent_backfill_tasks}
	MaxConcurrentBackfillTasks *float64 `field:"optional" json:"maxConcurrentBackfillTasks" yaml:"maxConcurrentBackfillTasks"`
	// Max concurrent CDC tasks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#max_concurrent_cdc_tasks DatastreamStream#max_concurrent_cdc_tasks}
	MaxConcurrentCdcTasks *float64 `field:"optional" json:"maxConcurrentCdcTasks" yaml:"maxConcurrentCdcTasks"`
	// The RPC priority to use for Spanner queries. Possible values: ["LOW", "MEDIUM", "HIGH"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#spanner_rpc_priority DatastreamStream#spanner_rpc_priority}
	SpannerRpcPriority *string `field:"optional" json:"spannerRpcPriority" yaml:"spannerRpcPriority"`
}

