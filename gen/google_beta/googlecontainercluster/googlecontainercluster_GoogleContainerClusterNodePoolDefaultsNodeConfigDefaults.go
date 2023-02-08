package googlecontainercluster


type GoogleContainerClusterNodePoolDefaultsNodeConfigDefaults struct {
	// gcfs_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#gcfs_config GoogleContainerCluster#gcfs_config}
	GcfsConfig *GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsGcfsConfig `field:"optional" json:"gcfsConfig" yaml:"gcfsConfig"`
}

