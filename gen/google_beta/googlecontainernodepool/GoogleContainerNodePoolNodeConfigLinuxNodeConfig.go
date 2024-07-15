package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigLinuxNodeConfig struct {
	// cgroupMode specifies the cgroup mode to be used on the node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_node_pool#cgroup_mode GoogleContainerNodePool#cgroup_mode}
	CgroupMode *string `field:"optional" json:"cgroupMode" yaml:"cgroupMode"`
	// The Linux kernel parameters to be applied to the nodes and all pods running on the nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_node_pool#sysctls GoogleContainerNodePool#sysctls}
	Sysctls *map[string]*string `field:"optional" json:"sysctls" yaml:"sysctls"`
}

