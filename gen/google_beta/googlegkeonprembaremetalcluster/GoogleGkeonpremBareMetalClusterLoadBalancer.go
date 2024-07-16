package googlegkeonprembaremetalcluster


type GoogleGkeonpremBareMetalClusterLoadBalancer struct {
	// port_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#port_config GoogleGkeonpremBareMetalCluster#port_config}
	PortConfig *GoogleGkeonpremBareMetalClusterLoadBalancerPortConfig `field:"required" json:"portConfig" yaml:"portConfig"`
	// vip_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#vip_config GoogleGkeonpremBareMetalCluster#vip_config}
	VipConfig *GoogleGkeonpremBareMetalClusterLoadBalancerVipConfig `field:"required" json:"vipConfig" yaml:"vipConfig"`
	// bgp_lb_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#bgp_lb_config GoogleGkeonpremBareMetalCluster#bgp_lb_config}
	BgpLbConfig *GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfig `field:"optional" json:"bgpLbConfig" yaml:"bgpLbConfig"`
	// manual_lb_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#manual_lb_config GoogleGkeonpremBareMetalCluster#manual_lb_config}
	ManualLbConfig *GoogleGkeonpremBareMetalClusterLoadBalancerManualLbConfig `field:"optional" json:"manualLbConfig" yaml:"manualLbConfig"`
	// metal_lb_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#metal_lb_config GoogleGkeonpremBareMetalCluster#metal_lb_config}
	MetalLbConfig *GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfig `field:"optional" json:"metalLbConfig" yaml:"metalLbConfig"`
}

