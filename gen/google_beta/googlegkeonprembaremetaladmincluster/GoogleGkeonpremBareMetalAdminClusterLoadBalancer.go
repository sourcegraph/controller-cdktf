package googlegkeonprembaremetaladmincluster


type GoogleGkeonpremBareMetalAdminClusterLoadBalancer struct {
	// port_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gkeonprem_bare_metal_admin_cluster#port_config GoogleGkeonpremBareMetalAdminCluster#port_config}
	PortConfig *GoogleGkeonpremBareMetalAdminClusterLoadBalancerPortConfig `field:"required" json:"portConfig" yaml:"portConfig"`
	// vip_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gkeonprem_bare_metal_admin_cluster#vip_config GoogleGkeonpremBareMetalAdminCluster#vip_config}
	VipConfig *GoogleGkeonpremBareMetalAdminClusterLoadBalancerVipConfig `field:"required" json:"vipConfig" yaml:"vipConfig"`
	// bgp_lb_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gkeonprem_bare_metal_admin_cluster#bgp_lb_config GoogleGkeonpremBareMetalAdminCluster#bgp_lb_config}
	BgpLbConfig *GoogleGkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfig `field:"optional" json:"bgpLbConfig" yaml:"bgpLbConfig"`
	// manual_lb_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gkeonprem_bare_metal_admin_cluster#manual_lb_config GoogleGkeonpremBareMetalAdminCluster#manual_lb_config}
	ManualLbConfig *GoogleGkeonpremBareMetalAdminClusterLoadBalancerManualLbConfig `field:"optional" json:"manualLbConfig" yaml:"manualLbConfig"`
}

