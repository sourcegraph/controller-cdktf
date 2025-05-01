package googlegkeonpremvmwareadmincluster


type GoogleGkeonpremVmwareAdminClusterLoadBalancer struct {
	// vip_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gkeonprem_vmware_admin_cluster#vip_config GoogleGkeonpremVmwareAdminCluster#vip_config}
	VipConfig *GoogleGkeonpremVmwareAdminClusterLoadBalancerVipConfig `field:"required" json:"vipConfig" yaml:"vipConfig"`
	// f5_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gkeonprem_vmware_admin_cluster#f5_config GoogleGkeonpremVmwareAdminCluster#f5_config}
	F5Config *GoogleGkeonpremVmwareAdminClusterLoadBalancerF5Config `field:"optional" json:"f5Config" yaml:"f5Config"`
	// manual_lb_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gkeonprem_vmware_admin_cluster#manual_lb_config GoogleGkeonpremVmwareAdminCluster#manual_lb_config}
	ManualLbConfig *GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfig `field:"optional" json:"manualLbConfig" yaml:"manualLbConfig"`
	// metal_lb_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gkeonprem_vmware_admin_cluster#metal_lb_config GoogleGkeonpremVmwareAdminCluster#metal_lb_config}
	MetalLbConfig *GoogleGkeonpremVmwareAdminClusterLoadBalancerMetalLbConfig `field:"optional" json:"metalLbConfig" yaml:"metalLbConfig"`
}

