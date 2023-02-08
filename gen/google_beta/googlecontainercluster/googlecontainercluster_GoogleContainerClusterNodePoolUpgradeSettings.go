package googlecontainercluster


type GoogleContainerClusterNodePoolUpgradeSettings struct {
	// The number of additional nodes that can be added to the node pool during an upgrade.
	//
	// Increasing max_surge raises the number of nodes that can be upgraded simultaneously. Can be set to 0 or greater.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#max_surge GoogleContainerCluster#max_surge}
	MaxSurge *float64 `field:"required" json:"maxSurge" yaml:"maxSurge"`
	// The number of nodes that can be simultaneously unavailable during an upgrade.
	//
	// Increasing max_unavailable raises the number of nodes that can be upgraded in parallel. Can be set to 0 or greater.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#max_unavailable GoogleContainerCluster#max_unavailable}
	MaxUnavailable *float64 `field:"required" json:"maxUnavailable" yaml:"maxUnavailable"`
}

