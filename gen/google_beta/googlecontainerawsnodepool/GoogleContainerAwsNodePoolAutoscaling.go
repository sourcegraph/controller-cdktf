package googlecontainerawsnodepool


type GoogleContainerAwsNodePoolAutoscaling struct {
	// Maximum number of nodes in the NodePool. Must be >= min_node_count.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_aws_node_pool#max_node_count GoogleContainerAwsNodePool#max_node_count}
	MaxNodeCount *float64 `field:"required" json:"maxNodeCount" yaml:"maxNodeCount"`
	// Minimum number of nodes in the NodePool. Must be >= 1 and <= max_node_count.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_aws_node_pool#min_node_count GoogleContainerAwsNodePool#min_node_count}
	MinNodeCount *float64 `field:"required" json:"minNodeCount" yaml:"minNodeCount"`
}

