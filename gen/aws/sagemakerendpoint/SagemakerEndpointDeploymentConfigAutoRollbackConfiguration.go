package sagemakerendpoint


type SagemakerEndpointDeploymentConfigAutoRollbackConfiguration struct {
	// alarms block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/sagemaker_endpoint#alarms SagemakerEndpoint#alarms}
	Alarms interface{} `field:"optional" json:"alarms" yaml:"alarms"`
}

