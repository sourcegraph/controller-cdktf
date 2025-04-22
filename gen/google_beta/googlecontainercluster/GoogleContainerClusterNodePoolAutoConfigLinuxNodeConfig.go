package googlecontainercluster


type GoogleContainerClusterNodePoolAutoConfigLinuxNodeConfig struct {
	// cgroupMode specifies the cgroup mode to be used on the node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_container_cluster#cgroup_mode GoogleContainerCluster#cgroup_mode}
	CgroupMode *string `field:"optional" json:"cgroupMode" yaml:"cgroupMode"`
}

