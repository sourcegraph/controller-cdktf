package loadbalancer


type LoadBalancerRandomSteering struct {
	// The default weight for pools in the load balancer that are not specified in the [`pool_weights`](#pool_weights) map.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#default_weight LoadBalancer#default_weight}
	DefaultWeight *float64 `field:"optional" json:"defaultWeight" yaml:"defaultWeight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools in the load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#pool_weights LoadBalancer#pool_weights}
	PoolWeights *map[string]*float64 `field:"optional" json:"poolWeights" yaml:"poolWeights"`
}

