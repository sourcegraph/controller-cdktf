package googleclouddeploydeliverypipeline


type GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentAnalysis struct {
	// Required. Duration of the analysis.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_clouddeploy_delivery_pipeline#duration GoogleClouddeployDeliveryPipeline#duration}
	Duration *string `field:"required" json:"duration" yaml:"duration"`
	// custom_checks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_clouddeploy_delivery_pipeline#custom_checks GoogleClouddeployDeliveryPipeline#custom_checks}
	CustomChecks interface{} `field:"optional" json:"customChecks" yaml:"customChecks"`
	// google_cloud block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_clouddeploy_delivery_pipeline#google_cloud GoogleClouddeployDeliveryPipeline#google_cloud}
	GoogleCloud *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentAnalysisGoogleCloud `field:"optional" json:"googleCloud" yaml:"googleCloud"`
}

