package containernodepool


type ContainerNodePoolNodeConfigLinuxNodeConfigNodeKernelModuleLoading struct {
	// The policy for kernel module loading.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_node_pool#policy ContainerNodePool#policy}
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
}

