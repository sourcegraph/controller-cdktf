package googlegkeonprembaremetaladmincluster


type GoogleGkeonpremBareMetalAdminClusterLoadBalancerVipConfig struct {
	// The VIP which you previously set aside for the Kubernetes API of this Bare Metal Admin Cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_admin_cluster#control_plane_vip GoogleGkeonpremBareMetalAdminCluster#control_plane_vip}
	ControlPlaneVip *string `field:"required" json:"controlPlaneVip" yaml:"controlPlaneVip"`
}

