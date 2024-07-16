package googlegkeonprembaremetalcluster


type GoogleGkeonpremBareMetalClusterControlPlane struct {
	// control_plane_node_pool_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#control_plane_node_pool_config GoogleGkeonpremBareMetalCluster#control_plane_node_pool_config}
	ControlPlaneNodePoolConfig *GoogleGkeonpremBareMetalClusterControlPlaneControlPlaneNodePoolConfig `field:"required" json:"controlPlaneNodePoolConfig" yaml:"controlPlaneNodePoolConfig"`
	// api_server_args block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#api_server_args GoogleGkeonpremBareMetalCluster#api_server_args}
	ApiServerArgs interface{} `field:"optional" json:"apiServerArgs" yaml:"apiServerArgs"`
}

