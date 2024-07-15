package googlegkeonprembaremetaladmincluster


type GoogleGkeonpremBareMetalAdminClusterStorageLvpShareConfigLvpConfig struct {
	// The host machine path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_admin_cluster#path GoogleGkeonpremBareMetalAdminCluster#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// The StorageClass name that PVs will be created with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_admin_cluster#storage_class GoogleGkeonpremBareMetalAdminCluster#storage_class}
	StorageClass *string `field:"required" json:"storageClass" yaml:"storageClass"`
}

