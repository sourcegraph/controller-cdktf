package googlegkeonprembaremetaladmincluster


type GoogleGkeonpremBareMetalAdminClusterControlPlane struct {
	// control_plane_node_pool_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_admin_cluster#control_plane_node_pool_config GoogleGkeonpremBareMetalAdminCluster#control_plane_node_pool_config}
	ControlPlaneNodePoolConfig *GoogleGkeonpremBareMetalAdminClusterControlPlaneControlPlaneNodePoolConfig `field:"required" json:"controlPlaneNodePoolConfig" yaml:"controlPlaneNodePoolConfig"`
	// api_server_args block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_admin_cluster#api_server_args GoogleGkeonpremBareMetalAdminCluster#api_server_args}
	ApiServerArgs interface{} `field:"optional" json:"apiServerArgs" yaml:"apiServerArgs"`
}

