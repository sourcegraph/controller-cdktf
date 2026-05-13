package googlecontainercluster


type GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfig struct {
	// accurate_time_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#accurate_time_config GoogleContainerCluster#accurate_time_config}
	AccurateTimeConfig *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigAccurateTimeConfig `field:"optional" json:"accurateTimeConfig" yaml:"accurateTimeConfig"`
	// cgroupMode specifies the cgroup mode to be used on the node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#cgroup_mode GoogleContainerCluster#cgroup_mode}
	CgroupMode *string `field:"optional" json:"cgroupMode" yaml:"cgroupMode"`
	// hugepages_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#hugepages_config GoogleContainerCluster#hugepages_config}
	HugepagesConfig *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigHugepagesConfig `field:"optional" json:"hugepagesConfig" yaml:"hugepagesConfig"`
	// node_kernel_module_loading block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#node_kernel_module_loading GoogleContainerCluster#node_kernel_module_loading}
	NodeKernelModuleLoading *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigNodeKernelModuleLoading `field:"optional" json:"nodeKernelModuleLoading" yaml:"nodeKernelModuleLoading"`
	// swap_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#swap_config GoogleContainerCluster#swap_config}
	SwapConfig *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfig `field:"optional" json:"swapConfig" yaml:"swapConfig"`
	// The Linux kernel parameters to be applied to the nodes and all pods running on the nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#sysctls GoogleContainerCluster#sysctls}
	Sysctls *map[string]*string `field:"optional" json:"sysctls" yaml:"sysctls"`
	// The Linux kernel transparent hugepage defrag setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#transparent_hugepage_defrag GoogleContainerCluster#transparent_hugepage_defrag}
	TransparentHugepageDefrag *string `field:"optional" json:"transparentHugepageDefrag" yaml:"transparentHugepageDefrag"`
	// The Linux kernel transparent hugepage setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#transparent_hugepage_enabled GoogleContainerCluster#transparent_hugepage_enabled}
	TransparentHugepageEnabled *string `field:"optional" json:"transparentHugepageEnabled" yaml:"transparentHugepageEnabled"`
}

