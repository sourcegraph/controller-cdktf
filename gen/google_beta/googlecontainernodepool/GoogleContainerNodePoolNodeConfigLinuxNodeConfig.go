package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigLinuxNodeConfig struct {
	// accurate_time_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#accurate_time_config GoogleContainerNodePool#accurate_time_config}
	AccurateTimeConfig *GoogleContainerNodePoolNodeConfigLinuxNodeConfigAccurateTimeConfig `field:"optional" json:"accurateTimeConfig" yaml:"accurateTimeConfig"`
	// cgroupMode specifies the cgroup mode to be used on the node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#cgroup_mode GoogleContainerNodePool#cgroup_mode}
	CgroupMode *string `field:"optional" json:"cgroupMode" yaml:"cgroupMode"`
	// hugepages_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#hugepages_config GoogleContainerNodePool#hugepages_config}
	HugepagesConfig *GoogleContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfig `field:"optional" json:"hugepagesConfig" yaml:"hugepagesConfig"`
	// node_kernel_module_loading block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#node_kernel_module_loading GoogleContainerNodePool#node_kernel_module_loading}
	NodeKernelModuleLoading *GoogleContainerNodePoolNodeConfigLinuxNodeConfigNodeKernelModuleLoading `field:"optional" json:"nodeKernelModuleLoading" yaml:"nodeKernelModuleLoading"`
	// swap_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#swap_config GoogleContainerNodePool#swap_config}
	SwapConfig *GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfig `field:"optional" json:"swapConfig" yaml:"swapConfig"`
	// The Linux kernel parameters to be applied to the nodes and all pods running on the nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#sysctls GoogleContainerNodePool#sysctls}
	Sysctls *map[string]*string `field:"optional" json:"sysctls" yaml:"sysctls"`
	// The Linux kernel transparent hugepage defrag setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#transparent_hugepage_defrag GoogleContainerNodePool#transparent_hugepage_defrag}
	TransparentHugepageDefrag *string `field:"optional" json:"transparentHugepageDefrag" yaml:"transparentHugepageDefrag"`
	// The Linux kernel transparent hugepage setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#transparent_hugepage_enabled GoogleContainerNodePool#transparent_hugepage_enabled}
	TransparentHugepageEnabled *string `field:"optional" json:"transparentHugepageEnabled" yaml:"transparentHugepageEnabled"`
}

