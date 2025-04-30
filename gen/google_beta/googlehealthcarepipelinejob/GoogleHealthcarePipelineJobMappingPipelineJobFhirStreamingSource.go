package googlehealthcarepipelinejob


type GoogleHealthcarePipelineJobMappingPipelineJobFhirStreamingSource struct {
	// The path to the FHIR store in the format projects/{projectId}/locations/{locationId}/datasets/{datasetId}/fhirStores/{fhirStoreId}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_healthcare_pipeline_job#fhir_store GoogleHealthcarePipelineJob#fhir_store}
	FhirStore *string `field:"required" json:"fhirStore" yaml:"fhirStore"`
	// Describes the streaming FHIR data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_healthcare_pipeline_job#description GoogleHealthcarePipelineJob#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

