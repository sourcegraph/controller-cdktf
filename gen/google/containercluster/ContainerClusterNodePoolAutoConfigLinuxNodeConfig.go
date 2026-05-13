package containercluster


type ContainerClusterNodePoolAutoConfigLinuxNodeConfig struct {
	// cgroupMode specifies the cgroup mode to be used on the node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#cgroup_mode ContainerCluster#cgroup_mode}
	CgroupMode *string `field:"optional" json:"cgroupMode" yaml:"cgroupMode"`
	// node_kernel_module_loading block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#node_kernel_module_loading ContainerCluster#node_kernel_module_loading}
	NodeKernelModuleLoading *ContainerClusterNodePoolAutoConfigLinuxNodeConfigNodeKernelModuleLoading `field:"optional" json:"nodeKernelModuleLoading" yaml:"nodeKernelModuleLoading"`
}

