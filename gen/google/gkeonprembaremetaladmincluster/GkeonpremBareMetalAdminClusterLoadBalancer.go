package gkeonprembaremetaladmincluster


type GkeonpremBareMetalAdminClusterLoadBalancer struct {
	// port_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/gkeonprem_bare_metal_admin_cluster#port_config GkeonpremBareMetalAdminCluster#port_config}
	PortConfig *GkeonpremBareMetalAdminClusterLoadBalancerPortConfig `field:"required" json:"portConfig" yaml:"portConfig"`
	// vip_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/gkeonprem_bare_metal_admin_cluster#vip_config GkeonpremBareMetalAdminCluster#vip_config}
	VipConfig *GkeonpremBareMetalAdminClusterLoadBalancerVipConfig `field:"required" json:"vipConfig" yaml:"vipConfig"`
	// bgp_lb_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/gkeonprem_bare_metal_admin_cluster#bgp_lb_config GkeonpremBareMetalAdminCluster#bgp_lb_config}
	BgpLbConfig *GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfig `field:"optional" json:"bgpLbConfig" yaml:"bgpLbConfig"`
	// manual_lb_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/gkeonprem_bare_metal_admin_cluster#manual_lb_config GkeonpremBareMetalAdminCluster#manual_lb_config}
	ManualLbConfig *GkeonpremBareMetalAdminClusterLoadBalancerManualLbConfig `field:"optional" json:"manualLbConfig" yaml:"manualLbConfig"`
}

