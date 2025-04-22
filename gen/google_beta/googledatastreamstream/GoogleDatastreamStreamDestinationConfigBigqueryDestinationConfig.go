package googledatastreamstream


type GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfig struct {
	// append_only block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_datastream_stream#append_only GoogleDatastreamStream#append_only}
	AppendOnly *GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigAppendOnly `field:"optional" json:"appendOnly" yaml:"appendOnly"`
	// blmt_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_datastream_stream#blmt_config GoogleDatastreamStream#blmt_config}
	BlmtConfig *GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigBlmtConfig `field:"optional" json:"blmtConfig" yaml:"blmtConfig"`
	// The guaranteed data freshness (in seconds) when querying tables created by the stream.
	//
	// Editing this field will only affect new tables created in the future, but existing tables
	// will not be impacted. Lower values mean that queries will return fresher data, but may result in higher cost.
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s". Defaults to 900s.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_datastream_stream#data_freshness GoogleDatastreamStream#data_freshness}
	DataFreshness *string `field:"optional" json:"dataFreshness" yaml:"dataFreshness"`
	// merge block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_datastream_stream#merge GoogleDatastreamStream#merge}
	Merge *GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigMerge `field:"optional" json:"merge" yaml:"merge"`
	// single_target_dataset block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_datastream_stream#single_target_dataset GoogleDatastreamStream#single_target_dataset}
	SingleTargetDataset *GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigSingleTargetDataset `field:"optional" json:"singleTargetDataset" yaml:"singleTargetDataset"`
	// source_hierarchy_datasets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_datastream_stream#source_hierarchy_datasets GoogleDatastreamStream#source_hierarchy_datasets}
	SourceHierarchyDatasets *GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasets `field:"optional" json:"sourceHierarchyDatasets" yaml:"sourceHierarchyDatasets"`
}

