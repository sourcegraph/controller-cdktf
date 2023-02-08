package googlecontainercluster


type GoogleContainerClusterNodeConfigWorkloadMetadataConfig struct {
	// Mode is the configuration for how to expose metadata to workloads running on the node.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#mode GoogleContainerCluster#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}

