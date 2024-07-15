package googledatastreamstream


type GoogleDatastreamStreamSourceConfigSqlServerSourceConfig struct {
	// exclude_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_stream#exclude_objects GoogleDatastreamStream#exclude_objects}
	ExcludeObjects *GoogleDatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjects `field:"optional" json:"excludeObjects" yaml:"excludeObjects"`
	// include_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_stream#include_objects GoogleDatastreamStream#include_objects}
	IncludeObjects *GoogleDatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjects `field:"optional" json:"includeObjects" yaml:"includeObjects"`
	// Max concurrent backfill tasks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_stream#max_concurrent_backfill_tasks GoogleDatastreamStream#max_concurrent_backfill_tasks}
	MaxConcurrentBackfillTasks *float64 `field:"optional" json:"maxConcurrentBackfillTasks" yaml:"maxConcurrentBackfillTasks"`
	// Max concurrent CDC tasks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_stream#max_concurrent_cdc_tasks GoogleDatastreamStream#max_concurrent_cdc_tasks}
	MaxConcurrentCdcTasks *float64 `field:"optional" json:"maxConcurrentCdcTasks" yaml:"maxConcurrentCdcTasks"`
}

