package deploymentv1


type DeploymentV1SpecTemplateSpecContainerEnvValueFromResourceFieldRef struct {
	// Resource to select.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/deployment_v1#resource DeploymentV1#resource}
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/deployment_v1#container_name DeploymentV1#container_name}.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/deployment_v1#divisor DeploymentV1#divisor}.
	Divisor *string `field:"optional" json:"divisor" yaml:"divisor"`
}

