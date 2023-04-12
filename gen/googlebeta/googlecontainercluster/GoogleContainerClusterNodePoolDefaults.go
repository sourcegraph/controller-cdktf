package googlecontainercluster


type GoogleContainerClusterNodePoolDefaults struct {
	// node_config_defaults block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#node_config_defaults GoogleContainerCluster#node_config_defaults}
	NodeConfigDefaults *GoogleContainerClusterNodePoolDefaultsNodeConfigDefaults `field:"optional" json:"nodeConfigDefaults" yaml:"nodeConfigDefaults"`
}

