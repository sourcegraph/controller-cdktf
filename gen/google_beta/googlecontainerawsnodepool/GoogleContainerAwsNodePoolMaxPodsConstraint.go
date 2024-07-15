package googlecontainerawsnodepool


type GoogleContainerAwsNodePoolMaxPodsConstraint struct {
	// The maximum number of pods to schedule on a single node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_aws_node_pool#max_pods_per_node GoogleContainerAwsNodePool#max_pods_per_node}
	MaxPodsPerNode *float64 `field:"required" json:"maxPodsPerNode" yaml:"maxPodsPerNode"`
}

