package loadbalancerpool


type LoadBalancerPoolOriginSteering struct {
	// Origin steering policy to be used. Available values: `""`, `hash`, `random`. Defaults to `random`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer_pool#policy LoadBalancerPool#policy}
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
}

