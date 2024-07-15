package googledatapipelinepipeline


type GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameter struct {
	// The job name to use for the created job.
	//
	// For an update job request, the job name should be the same as the existing running job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_pipeline_pipeline#job_name GoogleDataPipelinePipeline#job_name}
	JobName *string `field:"required" json:"jobName" yaml:"jobName"`
	// Cloud Storage path to a file with a JSON-serialized ContainerSpec as content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_pipeline_pipeline#container_spec_gcs_path GoogleDataPipelinePipeline#container_spec_gcs_path}
	ContainerSpecGcsPath *string `field:"optional" json:"containerSpecGcsPath" yaml:"containerSpecGcsPath"`
	// environment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_pipeline_pipeline#environment GoogleDataPipelinePipeline#environment}
	Environment *GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironment `field:"optional" json:"environment" yaml:"environment"`
	// Launch options for this Flex Template job.
	//
	// This is a common set of options across languages and templates. This should not be used to pass job parameters.
	// 'An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_pipeline_pipeline#launch_options GoogleDataPipelinePipeline#launch_options}
	LaunchOptions *map[string]*string `field:"optional" json:"launchOptions" yaml:"launchOptions"`
	// 'The parameters for the Flex Template.
	//
	// Example: {"numWorkers":"5"}'
	// 'An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_pipeline_pipeline#parameters GoogleDataPipelinePipeline#parameters}
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// 'Use this to pass transform name mappings for streaming update jobs.
	//
	// Example: {"oldTransformName":"newTransformName",...}'
	// 'An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_pipeline_pipeline#transform_name_mappings GoogleDataPipelinePipeline#transform_name_mappings}
	TransformNameMappings *map[string]*string `field:"optional" json:"transformNameMappings" yaml:"transformNameMappings"`
	// Set this to true if you are sending a request to update a running streaming job.
	//
	// When set, the job name should be the same as the running job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_pipeline_pipeline#update GoogleDataPipelinePipeline#update}
	Update interface{} `field:"optional" json:"update" yaml:"update"`
}

