package datacloudflareloadbalancerpools


type DataCloudflareLoadBalancerPoolsFilter struct {
	// A regular expression matching the name of the Load Balancer pool to lookup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/data-sources/load_balancer_pools#name DataCloudflareLoadBalancerPools#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

