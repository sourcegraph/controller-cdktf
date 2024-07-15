package googledatapipelinepipeline


type GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequest struct {
	// launch_parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_pipeline_pipeline#launch_parameter GoogleDataPipelinePipeline#launch_parameter}
	LaunchParameter *GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameter `field:"required" json:"launchParameter" yaml:"launchParameter"`
	// The regional endpoint to which to direct the request. For example, us-central1, us-west1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_pipeline_pipeline#location GoogleDataPipelinePipeline#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The ID of the Cloud Platform project that the job belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_pipeline_pipeline#project_id GoogleDataPipelinePipeline#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// If true, the request is validated but not actually executed. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_pipeline_pipeline#validate_only GoogleDataPipelinePipeline#validate_only}
	ValidateOnly interface{} `field:"optional" json:"validateOnly" yaml:"validateOnly"`
}

