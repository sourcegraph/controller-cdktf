package googlegkeonprembaremetaladmincluster


type GoogleGkeonpremBareMetalAdminClusterStorageLvpShareConfig struct {
	// lvp_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_admin_cluster#lvp_config GoogleGkeonpremBareMetalAdminCluster#lvp_config}
	LvpConfig *GoogleGkeonpremBareMetalAdminClusterStorageLvpShareConfigLvpConfig `field:"required" json:"lvpConfig" yaml:"lvpConfig"`
	// The number of subdirectories to create under path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_admin_cluster#shared_path_pv_count GoogleGkeonpremBareMetalAdminCluster#shared_path_pv_count}
	SharedPathPvCount *float64 `field:"optional" json:"sharedPathPvCount" yaml:"sharedPathPvCount"`
}

