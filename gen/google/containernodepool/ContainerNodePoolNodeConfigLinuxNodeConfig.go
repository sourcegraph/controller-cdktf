package containernodepool


type ContainerNodePoolNodeConfigLinuxNodeConfig struct {
	// accurate_time_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_node_pool#accurate_time_config ContainerNodePool#accurate_time_config}
	AccurateTimeConfig *ContainerNodePoolNodeConfigLinuxNodeConfigAccurateTimeConfig `field:"optional" json:"accurateTimeConfig" yaml:"accurateTimeConfig"`
	// cgroupMode specifies the cgroup mode to be used on the node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_node_pool#cgroup_mode ContainerNodePool#cgroup_mode}
	CgroupMode *string `field:"optional" json:"cgroupMode" yaml:"cgroupMode"`
	// hugepages_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_node_pool#hugepages_config ContainerNodePool#hugepages_config}
	HugepagesConfig *ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfig `field:"optional" json:"hugepagesConfig" yaml:"hugepagesConfig"`
	// node_kernel_module_loading block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_node_pool#node_kernel_module_loading ContainerNodePool#node_kernel_module_loading}
	NodeKernelModuleLoading *ContainerNodePoolNodeConfigLinuxNodeConfigNodeKernelModuleLoading `field:"optional" json:"nodeKernelModuleLoading" yaml:"nodeKernelModuleLoading"`
	// swap_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_node_pool#swap_config ContainerNodePool#swap_config}
	SwapConfig *ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfig `field:"optional" json:"swapConfig" yaml:"swapConfig"`
	// The Linux kernel parameters to be applied to the nodes and all pods running on the nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_node_pool#sysctls ContainerNodePool#sysctls}
	Sysctls *map[string]*string `field:"optional" json:"sysctls" yaml:"sysctls"`
	// The Linux kernel transparent hugepage defrag setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_node_pool#transparent_hugepage_defrag ContainerNodePool#transparent_hugepage_defrag}
	TransparentHugepageDefrag *string `field:"optional" json:"transparentHugepageDefrag" yaml:"transparentHugepageDefrag"`
	// The Linux kernel transparent hugepage setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_node_pool#transparent_hugepage_enabled ContainerNodePool#transparent_hugepage_enabled}
	TransparentHugepageEnabled *string `field:"optional" json:"transparentHugepageEnabled" yaml:"transparentHugepageEnabled"`
}

