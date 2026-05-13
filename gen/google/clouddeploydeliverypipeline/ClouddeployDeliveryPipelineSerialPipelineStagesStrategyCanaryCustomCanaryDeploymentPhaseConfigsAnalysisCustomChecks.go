package clouddeploydeliverypipeline


type ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsAnalysisCustomChecks struct {
	// Required. Unique identifier for the custom check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/clouddeploy_delivery_pipeline#id ClouddeployDeliveryPipeline#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Optional. Frequency of the custom check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/clouddeploy_delivery_pipeline#frequency ClouddeployDeliveryPipeline#frequency}
	Frequency *string `field:"optional" json:"frequency" yaml:"frequency"`
	// task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/clouddeploy_delivery_pipeline#task ClouddeployDeliveryPipeline#task}
	Task *ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsAnalysisCustomChecksTask `field:"optional" json:"task" yaml:"task"`
}

