package googlegkeonprembaremetaladmincluster


type GoogleGkeonpremBareMetalAdminClusterStorage struct {
	// lvp_node_mounts_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_admin_cluster#lvp_node_mounts_config GoogleGkeonpremBareMetalAdminCluster#lvp_node_mounts_config}
	LvpNodeMountsConfig *GoogleGkeonpremBareMetalAdminClusterStorageLvpNodeMountsConfig `field:"required" json:"lvpNodeMountsConfig" yaml:"lvpNodeMountsConfig"`
	// lvp_share_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_admin_cluster#lvp_share_config GoogleGkeonpremBareMetalAdminCluster#lvp_share_config}
	LvpShareConfig *GoogleGkeonpremBareMetalAdminClusterStorageLvpShareConfig `field:"required" json:"lvpShareConfig" yaml:"lvpShareConfig"`
}

