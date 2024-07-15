package googlecontainerawsnodepool


type GoogleContainerAwsNodePoolConfigSpotConfig struct {
	// List of AWS EC2 instance types for creating a spot node pool's nodes.
	//
	// The specified instance types must have the same number of CPUs and memory. You can use the Amazon EC2 Instance Selector tool (https://github.com/aws/amazon-ec2-instance-selector) to choose instance types with matching CPU and memory
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_aws_node_pool#instance_types GoogleContainerAwsNodePool#instance_types}
	InstanceTypes *[]*string `field:"required" json:"instanceTypes" yaml:"instanceTypes"`
}

