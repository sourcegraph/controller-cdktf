package googlecontainercluster


type GoogleContainerClusterAddonsConfigPodSnapshotConfig struct {
	// Whether the Pod Snapshot feature is enabled for this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#enabled GoogleContainerCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

