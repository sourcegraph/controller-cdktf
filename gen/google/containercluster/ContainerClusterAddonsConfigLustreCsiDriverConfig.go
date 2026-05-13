package containercluster


type ContainerClusterAddonsConfigLustreCsiDriverConfig struct {
	// Whether the Lustre CSI driver is enabled for this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#enabled ContainerCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// When set to true, this disables multi-NIC support for the Lustre CSI driver.
	//
	// By default, GKE enables multi-NIC support, which
	// 										allows the Lustre CSI driver to automatically detect and configure all suitable network interfaces on a node to maximize I/O performance for demanding workloads.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#disable_multi_nic ContainerCluster#disable_multi_nic}
	DisableMultiNic interface{} `field:"optional" json:"disableMultiNic" yaml:"disableMultiNic"`
	// If set to true, the Lustre CSI driver will initialize LNet (the virtual network layer for Lustre kernel module) using port 6988.
	//
	// This flag is required to workaround a port conflict with the gke-metadata-server on GKE nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#enable_legacy_lustre_port ContainerCluster#enable_legacy_lustre_port}
	EnableLegacyLustrePort interface{} `field:"optional" json:"enableLegacyLustrePort" yaml:"enableLegacyLustrePort"`
}

