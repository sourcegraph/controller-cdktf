package googlegkeonprembaremetaladmincluster


type GoogleGkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigAddressPools struct {
	// The addresses that are part of this pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gkeonprem_bare_metal_admin_cluster#addresses GoogleGkeonpremBareMetalAdminCluster#addresses}
	Addresses *[]*string `field:"optional" json:"addresses" yaml:"addresses"`
	// This avoids buggy consumer devices mistakenly dropping IPv4 traffic for those special IP addresses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gkeonprem_bare_metal_admin_cluster#avoid_buggy_ips GoogleGkeonpremBareMetalAdminCluster#avoid_buggy_ips}
	AvoidBuggyIps interface{} `field:"optional" json:"avoidBuggyIps" yaml:"avoidBuggyIps"`
	// If true, prevent IP addresses from being automatically assigned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gkeonprem_bare_metal_admin_cluster#manual_assign GoogleGkeonpremBareMetalAdminCluster#manual_assign}
	ManualAssign interface{} `field:"optional" json:"manualAssign" yaml:"manualAssign"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gkeonprem_bare_metal_admin_cluster#pool GoogleGkeonpremBareMetalAdminCluster#pool}.
	Pool *string `field:"optional" json:"pool" yaml:"pool"`
}

