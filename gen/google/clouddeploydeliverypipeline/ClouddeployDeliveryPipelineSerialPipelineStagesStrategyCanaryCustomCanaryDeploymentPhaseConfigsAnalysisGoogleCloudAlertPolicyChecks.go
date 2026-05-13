package clouddeploydeliverypipeline


type ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsAnalysisGoogleCloudAlertPolicyChecks struct {
	// Required. The list of alert policy names to check. Format: `projects/{project}/alertPolicies/{alert_policy}`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/clouddeploy_delivery_pipeline#alert_policies ClouddeployDeliveryPipeline#alert_policies}
	AlertPolicies *[]*string `field:"required" json:"alertPolicies" yaml:"alertPolicies"`
	// Required. Unique identifier for the alert policy check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/clouddeploy_delivery_pipeline#id ClouddeployDeliveryPipeline#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Optional. Labels to filter the alert policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/clouddeploy_delivery_pipeline#labels ClouddeployDeliveryPipeline#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
}

