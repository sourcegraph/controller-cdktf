package datacloudflareloadbalancerpools


type DataCloudflareLoadBalancerPoolsFilter struct {
	// A regular expression matching the name of the Load Balancer pool to lookup.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/d/load_balancer_pools#name DataCloudflareLoadBalancerPools#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

