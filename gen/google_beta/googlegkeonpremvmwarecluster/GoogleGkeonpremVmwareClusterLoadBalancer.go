package googlegkeonpremvmwarecluster


type GoogleGkeonpremVmwareClusterLoadBalancer struct {
	// f5_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_vmware_cluster#f5_config GoogleGkeonpremVmwareCluster#f5_config}
	F5Config *GoogleGkeonpremVmwareClusterLoadBalancerF5Config `field:"optional" json:"f5Config" yaml:"f5Config"`
	// manual_lb_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_vmware_cluster#manual_lb_config GoogleGkeonpremVmwareCluster#manual_lb_config}
	ManualLbConfig *GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfig `field:"optional" json:"manualLbConfig" yaml:"manualLbConfig"`
	// metal_lb_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_vmware_cluster#metal_lb_config GoogleGkeonpremVmwareCluster#metal_lb_config}
	MetalLbConfig *GoogleGkeonpremVmwareClusterLoadBalancerMetalLbConfig `field:"optional" json:"metalLbConfig" yaml:"metalLbConfig"`
	// vip_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_vmware_cluster#vip_config GoogleGkeonpremVmwareCluster#vip_config}
	VipConfig *GoogleGkeonpremVmwareClusterLoadBalancerVipConfig `field:"optional" json:"vipConfig" yaml:"vipConfig"`
}

