package googlecontainercluster


type GoogleContainerClusterWorkloadAltsConfig struct {
	// Whether the alts handshaker should be enabled or not for direct-path. Requires Workload Identity (workloadPool must be non-empty).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#enable_alts GoogleContainerCluster#enable_alts}
	EnableAlts interface{} `field:"required" json:"enableAlts" yaml:"enableAlts"`
}

