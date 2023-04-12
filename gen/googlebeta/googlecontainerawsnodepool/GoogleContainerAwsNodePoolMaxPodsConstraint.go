package googlecontainerawsnodepool


type GoogleContainerAwsNodePoolMaxPodsConstraint struct {
	// The maximum number of pods to schedule on a single node.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_aws_node_pool#max_pods_per_node GoogleContainerAwsNodePool#max_pods_per_node}
	MaxPodsPerNode *float64 `field:"required" json:"maxPodsPerNode" yaml:"maxPodsPerNode"`
}

