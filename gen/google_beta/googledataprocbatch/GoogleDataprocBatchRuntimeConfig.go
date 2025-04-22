package googledataprocbatch


type GoogleDataprocBatchRuntimeConfig struct {
	// autotuning_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dataproc_batch#autotuning_config GoogleDataprocBatch#autotuning_config}
	AutotuningConfig *GoogleDataprocBatchRuntimeConfigAutotuningConfig `field:"optional" json:"autotuningConfig" yaml:"autotuningConfig"`
	// Optional. Cohort identifier. Identifies families of the workloads having the same shape, e.g. daily ETL jobs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dataproc_batch#cohort GoogleDataprocBatch#cohort}
	Cohort *string `field:"optional" json:"cohort" yaml:"cohort"`
	// Optional custom container image for the job runtime environment. If not specified, a default container image will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dataproc_batch#container_image GoogleDataprocBatch#container_image}
	ContainerImage *string `field:"optional" json:"containerImage" yaml:"containerImage"`
	// A mapping of property names to values, which are used to configure workload execution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dataproc_batch#properties GoogleDataprocBatch#properties}
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// Version of the batch runtime.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dataproc_batch#version GoogleDataprocBatch#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

