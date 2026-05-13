package clouddeploydeliverypipeline


type ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsAnalysis struct {
	// Required. Duration of the analysis.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/clouddeploy_delivery_pipeline#duration ClouddeployDeliveryPipeline#duration}
	Duration *string `field:"required" json:"duration" yaml:"duration"`
	// custom_checks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/clouddeploy_delivery_pipeline#custom_checks ClouddeployDeliveryPipeline#custom_checks}
	CustomChecks interface{} `field:"optional" json:"customChecks" yaml:"customChecks"`
	// google_cloud block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/clouddeploy_delivery_pipeline#google_cloud ClouddeployDeliveryPipeline#google_cloud}
	GoogleCloud *ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsAnalysisGoogleCloud `field:"optional" json:"googleCloud" yaml:"googleCloud"`
}

