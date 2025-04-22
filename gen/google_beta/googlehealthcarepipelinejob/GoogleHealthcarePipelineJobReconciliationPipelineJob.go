package googlehealthcarepipelinejob


type GoogleHealthcarePipelineJobReconciliationPipelineJob struct {
	// Specifies the top level directory of the matching configs used in all mapping pipelines, which extract properties for resources to be matched on.
	//
	// Example: gs://{bucket-id}/{path/to/matching/configs}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_healthcare_pipeline_job#matching_uri_prefix GoogleHealthcarePipelineJob#matching_uri_prefix}
	MatchingUriPrefix *string `field:"required" json:"matchingUriPrefix" yaml:"matchingUriPrefix"`
	// merge_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_healthcare_pipeline_job#merge_config GoogleHealthcarePipelineJob#merge_config}
	MergeConfig *GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfig `field:"required" json:"mergeConfig" yaml:"mergeConfig"`
	// The harmonized FHIR store to write harmonized FHIR resources to, in the format of: project/{projectID}/locations/{locationID}/datasets/{datasetName}/fhirStores/{id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_healthcare_pipeline_job#fhir_store_destination GoogleHealthcarePipelineJob#fhir_store_destination}
	FhirStoreDestination *string `field:"optional" json:"fhirStoreDestination" yaml:"fhirStoreDestination"`
}

