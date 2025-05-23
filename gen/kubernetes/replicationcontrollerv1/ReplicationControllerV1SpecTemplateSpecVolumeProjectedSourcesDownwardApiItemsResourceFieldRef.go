package replicationcontrollerv1


type ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRef struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/replication_controller_v1#container_name ReplicationControllerV1#container_name}.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// Resource to select.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/replication_controller_v1#resource ReplicationControllerV1#resource}
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/replication_controller_v1#divisor ReplicationControllerV1#divisor}.
	Divisor *string `field:"optional" json:"divisor" yaml:"divisor"`
}

