package googlegkeonprembaremetalcluster


type GoogleGkeonpremBareMetalClusterStorage struct {
	// lvp_node_mounts_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#lvp_node_mounts_config GoogleGkeonpremBareMetalCluster#lvp_node_mounts_config}
	LvpNodeMountsConfig *GoogleGkeonpremBareMetalClusterStorageLvpNodeMountsConfig `field:"required" json:"lvpNodeMountsConfig" yaml:"lvpNodeMountsConfig"`
	// lvp_share_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#lvp_share_config GoogleGkeonpremBareMetalCluster#lvp_share_config}
	LvpShareConfig *GoogleGkeonpremBareMetalClusterStorageLvpShareConfig `field:"required" json:"lvpShareConfig" yaml:"lvpShareConfig"`
}

