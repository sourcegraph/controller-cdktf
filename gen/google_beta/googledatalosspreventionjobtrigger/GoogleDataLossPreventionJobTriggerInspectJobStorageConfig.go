package googledatalosspreventionjobtrigger


type GoogleDataLossPreventionJobTriggerInspectJobStorageConfig struct {
	// big_query_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#big_query_options GoogleDataLossPreventionJobTrigger#big_query_options}
	BigQueryOptions *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptions `field:"optional" json:"bigQueryOptions" yaml:"bigQueryOptions"`
	// cloud_storage_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#cloud_storage_options GoogleDataLossPreventionJobTrigger#cloud_storage_options}
	CloudStorageOptions *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptions `field:"optional" json:"cloudStorageOptions" yaml:"cloudStorageOptions"`
	// datastore_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#datastore_options GoogleDataLossPreventionJobTrigger#datastore_options}
	DatastoreOptions *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigDatastoreOptions `field:"optional" json:"datastoreOptions" yaml:"datastoreOptions"`
	// hybrid_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#hybrid_options GoogleDataLossPreventionJobTrigger#hybrid_options}
	HybridOptions *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigHybridOptions `field:"optional" json:"hybridOptions" yaml:"hybridOptions"`
	// timespan_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#timespan_config GoogleDataLossPreventionJobTrigger#timespan_config}
	TimespanConfig *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigTimespanConfig `field:"optional" json:"timespanConfig" yaml:"timespanConfig"`
}

