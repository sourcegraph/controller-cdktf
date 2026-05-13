package googledatastreamstream


type GoogleDatastreamStreamSourceConfigMongodbSourceConfig struct {
	// exclude_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_stream#exclude_objects GoogleDatastreamStream#exclude_objects}
	ExcludeObjects *GoogleDatastreamStreamSourceConfigMongodbSourceConfigExcludeObjects `field:"optional" json:"excludeObjects" yaml:"excludeObjects"`
	// include_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_stream#include_objects GoogleDatastreamStream#include_objects}
	IncludeObjects *GoogleDatastreamStreamSourceConfigMongodbSourceConfigIncludeObjects `field:"optional" json:"includeObjects" yaml:"includeObjects"`
	// Optional.
	//
	// Maximum number of concurrent backfill tasks. The number
	// should be non-negative and less than or equal to 50. If not set
	// (or set to 0), the system''s default value is used
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_stream#max_concurrent_backfill_tasks GoogleDatastreamStream#max_concurrent_backfill_tasks}
	MaxConcurrentBackfillTasks *float64 `field:"optional" json:"maxConcurrentBackfillTasks" yaml:"maxConcurrentBackfillTasks"`
}

