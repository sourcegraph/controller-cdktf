package googlecontainercluster


type GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfig struct {
	// The Linux kernel parameters to be applied to the nodes and all pods running on the nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_container_cluster#sysctls GoogleContainerCluster#sysctls}
	Sysctls *map[string]*string `field:"required" json:"sysctls" yaml:"sysctls"`
}

