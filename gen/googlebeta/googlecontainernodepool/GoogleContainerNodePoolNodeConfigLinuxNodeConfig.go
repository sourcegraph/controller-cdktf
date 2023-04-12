package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigLinuxNodeConfig struct {
	// The Linux kernel parameters to be applied to the nodes and all pods running on the nodes.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_node_pool#sysctls GoogleContainerNodePool#sysctls}
	Sysctls *map[string]*string `field:"required" json:"sysctls" yaml:"sysctls"`
}

