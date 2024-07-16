package googlecontainercluster


type GoogleContainerClusterNodePoolDefaultsNodeConfigDefaults struct {
	// containerd_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#containerd_config GoogleContainerCluster#containerd_config}
	ContainerdConfig *GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfig `field:"optional" json:"containerdConfig" yaml:"containerdConfig"`
	// gcfs_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#gcfs_config GoogleContainerCluster#gcfs_config}
	GcfsConfig *GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsGcfsConfig `field:"optional" json:"gcfsConfig" yaml:"gcfsConfig"`
	// Type of logging agent that is used as the default value for node pools in the cluster.
	//
	// Valid values include DEFAULT and MAX_THROUGHPUT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#logging_variant GoogleContainerCluster#logging_variant}
	LoggingVariant *string `field:"optional" json:"loggingVariant" yaml:"loggingVariant"`
}

