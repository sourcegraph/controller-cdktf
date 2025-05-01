package googlegkeonpremvmwareadmincluster


type GoogleGkeonpremVmwareAdminClusterLoadBalancerVipConfig struct {
	// The VIP which you previously set aside for the Kubernetes API of this VMware Admin Cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gkeonprem_vmware_admin_cluster#control_plane_vip GoogleGkeonpremVmwareAdminCluster#control_plane_vip}
	ControlPlaneVip *string `field:"required" json:"controlPlaneVip" yaml:"controlPlaneVip"`
	// The VIP to configure the load balancer for add-ons.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gkeonprem_vmware_admin_cluster#addons_vip GoogleGkeonpremVmwareAdminCluster#addons_vip}
	AddonsVip *string `field:"optional" json:"addonsVip" yaml:"addonsVip"`
}

