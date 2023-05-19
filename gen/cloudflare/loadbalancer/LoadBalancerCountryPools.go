package loadbalancer


type LoadBalancerCountryPools struct {
	// A country code which can be determined with the Load Balancing Regions API described [here](https://developers.cloudflare.com/load-balancing/reference/region-mapping-api/). Multiple entries should not be specified with the same country.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#country LoadBalancer#country}
	Country *string `field:"required" json:"country" yaml:"country"`
	// A list of pool IDs in failover priority to use in the given country.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#pool_ids LoadBalancer#pool_ids}
	PoolIds *[]*string `field:"required" json:"poolIds" yaml:"poolIds"`
}

