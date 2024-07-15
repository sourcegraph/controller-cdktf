package googlegkeonprembaremetalcluster


type GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPools struct {
	// The addresses that are part of this pool.
	//
	// Each address must be either in the CIDR form (1.2.3.0/24) or range form (1.2.3.1-1.2.3.5).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#addresses GoogleGkeonpremBareMetalCluster#addresses}
	Addresses *[]*string `field:"required" json:"addresses" yaml:"addresses"`
	// The name of the address pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#pool GoogleGkeonpremBareMetalCluster#pool}
	Pool *string `field:"required" json:"pool" yaml:"pool"`
	// If true, avoid using IPs ending in .0 or .255. This avoids buggy consumer devices mistakenly dropping IPv4 traffic for those special IP addresses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#avoid_buggy_ips GoogleGkeonpremBareMetalCluster#avoid_buggy_ips}
	AvoidBuggyIps interface{} `field:"optional" json:"avoidBuggyIps" yaml:"avoidBuggyIps"`
	// If true, prevent IP addresses from being automatically assigned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#manual_assign GoogleGkeonpremBareMetalCluster#manual_assign}
	ManualAssign *string `field:"optional" json:"manualAssign" yaml:"manualAssign"`
}

