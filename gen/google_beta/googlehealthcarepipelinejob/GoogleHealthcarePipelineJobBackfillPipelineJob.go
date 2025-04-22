package googlehealthcarepipelinejob


type GoogleHealthcarePipelineJobBackfillPipelineJob struct {
	// Specifies the mapping pipeline job to backfill, the name format should follow: projects/{projectId}/locations/{locationId}/datasets/{datasetId}/pipelineJobs/{pipelineJobId}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_healthcare_pipeline_job#mapping_pipeline_job GoogleHealthcarePipelineJob#mapping_pipeline_job}
	MappingPipelineJob *string `field:"optional" json:"mappingPipelineJob" yaml:"mappingPipelineJob"`
}

