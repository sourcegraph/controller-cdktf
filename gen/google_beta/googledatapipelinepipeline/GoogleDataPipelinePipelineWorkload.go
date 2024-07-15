package googledatapipelinepipeline


type GoogleDataPipelinePipelineWorkload struct {
	// dataflow_flex_template_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_pipeline_pipeline#dataflow_flex_template_request GoogleDataPipelinePipeline#dataflow_flex_template_request}
	DataflowFlexTemplateRequest *GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequest `field:"optional" json:"dataflowFlexTemplateRequest" yaml:"dataflowFlexTemplateRequest"`
	// dataflow_launch_template_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_pipeline_pipeline#dataflow_launch_template_request GoogleDataPipelinePipeline#dataflow_launch_template_request}
	DataflowLaunchTemplateRequest *GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequest `field:"optional" json:"dataflowLaunchTemplateRequest" yaml:"dataflowLaunchTemplateRequest"`
}

