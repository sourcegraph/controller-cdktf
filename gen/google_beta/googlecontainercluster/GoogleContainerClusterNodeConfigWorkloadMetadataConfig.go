package googlecontainercluster


type GoogleContainerClusterNodeConfigWorkloadMetadataConfig struct {
	// Mode is the configuration for how to expose metadata to workloads running on the node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#mode GoogleContainerCluster#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}

