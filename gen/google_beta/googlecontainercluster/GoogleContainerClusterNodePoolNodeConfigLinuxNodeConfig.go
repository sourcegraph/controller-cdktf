package googlecontainercluster


type GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfig struct {
	// cgroupMode specifies the cgroup mode to be used on the node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_container_cluster#cgroup_mode GoogleContainerCluster#cgroup_mode}
	CgroupMode *string `field:"optional" json:"cgroupMode" yaml:"cgroupMode"`
	// hugepages_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_container_cluster#hugepages_config GoogleContainerCluster#hugepages_config}
	HugepagesConfig *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigHugepagesConfig `field:"optional" json:"hugepagesConfig" yaml:"hugepagesConfig"`
	// The Linux kernel parameters to be applied to the nodes and all pods running on the nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_container_cluster#sysctls GoogleContainerCluster#sysctls}
	Sysctls *map[string]*string `field:"optional" json:"sysctls" yaml:"sysctls"`
}

