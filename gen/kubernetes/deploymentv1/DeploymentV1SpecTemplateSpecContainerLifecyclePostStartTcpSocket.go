package deploymentv1


type DeploymentV1SpecTemplateSpecContainerLifecyclePostStartTcpSocket struct {
	// Number or name of the port to access on the container.
	//
	// Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/deployment_v1#port DeploymentV1#port}
	Port *string `field:"required" json:"port" yaml:"port"`
}

