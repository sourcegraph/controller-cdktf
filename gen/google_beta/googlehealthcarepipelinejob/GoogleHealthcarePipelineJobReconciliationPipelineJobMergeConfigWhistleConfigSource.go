package googlehealthcarepipelinejob


type GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSource struct {
	// Directory path where all the Whistle files are located. Example: gs://{bucket-id}/{path/to/import-root/dir}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_healthcare_pipeline_job#import_uri_prefix GoogleHealthcarePipelineJob#import_uri_prefix}
	ImportUriPrefix *string `field:"required" json:"importUriPrefix" yaml:"importUriPrefix"`
	// Main configuration file which has the entrypoint or the root function. Example: gs://{bucket-id}/{path/to/import-root/dir}/entrypoint-file-name.wstl.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_healthcare_pipeline_job#uri GoogleHealthcarePipelineJob#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
}

