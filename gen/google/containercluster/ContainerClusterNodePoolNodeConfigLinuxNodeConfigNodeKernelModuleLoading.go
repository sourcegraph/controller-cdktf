package containercluster


type ContainerClusterNodePoolNodeConfigLinuxNodeConfigNodeKernelModuleLoading struct {
	// The policy for kernel module loading.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#policy ContainerCluster#policy}
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
}

