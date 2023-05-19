package loadbalancerpool


type LoadBalancerPoolLoadShedding struct {
	// Percent of traffic to shed 0 - 100. Defaults to `0`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer_pool#default_percent LoadBalancerPool#default_percent}
	DefaultPercent *float64 `field:"optional" json:"defaultPercent" yaml:"defaultPercent"`
	// Method of shedding traffic. Available values: `""`, `hash`, `random`. Defaults to `""`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer_pool#default_policy LoadBalancerPool#default_policy}
	DefaultPolicy *string `field:"optional" json:"defaultPolicy" yaml:"defaultPolicy"`
	// Percent of session traffic to shed 0 - 100. Defaults to `0`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer_pool#session_percent LoadBalancerPool#session_percent}
	SessionPercent *float64 `field:"optional" json:"sessionPercent" yaml:"sessionPercent"`
	// Method of shedding traffic. Available values: `""`, `hash`. Defaults to `""`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer_pool#session_policy LoadBalancerPool#session_policy}
	SessionPolicy *string `field:"optional" json:"sessionPolicy" yaml:"sessionPolicy"`
}

