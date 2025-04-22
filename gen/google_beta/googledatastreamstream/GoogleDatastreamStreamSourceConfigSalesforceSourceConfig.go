package googledatastreamstream


type GoogleDatastreamStreamSourceConfigSalesforceSourceConfig struct {
	// Salesforce objects polling interval.
	//
	// The interval at which new changes will be polled for each object. The duration must be between 5 minutes and 24 hours.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_datastream_stream#polling_interval GoogleDatastreamStream#polling_interval}
	PollingInterval *string `field:"required" json:"pollingInterval" yaml:"pollingInterval"`
	// exclude_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_datastream_stream#exclude_objects GoogleDatastreamStream#exclude_objects}
	ExcludeObjects *GoogleDatastreamStreamSourceConfigSalesforceSourceConfigExcludeObjects `field:"optional" json:"excludeObjects" yaml:"excludeObjects"`
	// include_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_datastream_stream#include_objects GoogleDatastreamStream#include_objects}
	IncludeObjects *GoogleDatastreamStreamSourceConfigSalesforceSourceConfigIncludeObjects `field:"optional" json:"includeObjects" yaml:"includeObjects"`
}

