package podv1


type PodV1SpecContainerEnvValueFromResourceFieldRef struct {
	// Resource to select.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod_v1#resource PodV1#resource}
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod_v1#container_name PodV1#container_name}.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod_v1#divisor PodV1#divisor}.
	Divisor *string `field:"optional" json:"divisor" yaml:"divisor"`
}

