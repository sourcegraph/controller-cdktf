package containercluster


type ContainerClusterNodeConfigLinuxNodeConfig struct {
	// accurate_time_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#accurate_time_config ContainerCluster#accurate_time_config}
	AccurateTimeConfig *ContainerClusterNodeConfigLinuxNodeConfigAccurateTimeConfig `field:"optional" json:"accurateTimeConfig" yaml:"accurateTimeConfig"`
	// cgroupMode specifies the cgroup mode to be used on the node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#cgroup_mode ContainerCluster#cgroup_mode}
	CgroupMode *string `field:"optional" json:"cgroupMode" yaml:"cgroupMode"`
	// hugepages_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#hugepages_config ContainerCluster#hugepages_config}
	HugepagesConfig *ContainerClusterNodeConfigLinuxNodeConfigHugepagesConfig `field:"optional" json:"hugepagesConfig" yaml:"hugepagesConfig"`
	// node_kernel_module_loading block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#node_kernel_module_loading ContainerCluster#node_kernel_module_loading}
	NodeKernelModuleLoading *ContainerClusterNodeConfigLinuxNodeConfigNodeKernelModuleLoading `field:"optional" json:"nodeKernelModuleLoading" yaml:"nodeKernelModuleLoading"`
	// swap_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#swap_config ContainerCluster#swap_config}
	SwapConfig *ContainerClusterNodeConfigLinuxNodeConfigSwapConfig `field:"optional" json:"swapConfig" yaml:"swapConfig"`
	// The Linux kernel parameters to be applied to the nodes and all pods running on the nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#sysctls ContainerCluster#sysctls}
	Sysctls *map[string]*string `field:"optional" json:"sysctls" yaml:"sysctls"`
	// The Linux kernel transparent hugepage defrag setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#transparent_hugepage_defrag ContainerCluster#transparent_hugepage_defrag}
	TransparentHugepageDefrag *string `field:"optional" json:"transparentHugepageDefrag" yaml:"transparentHugepageDefrag"`
	// The Linux kernel transparent hugepage setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#transparent_hugepage_enabled ContainerCluster#transparent_hugepage_enabled}
	TransparentHugepageEnabled *string `field:"optional" json:"transparentHugepageEnabled" yaml:"transparentHugepageEnabled"`
}

