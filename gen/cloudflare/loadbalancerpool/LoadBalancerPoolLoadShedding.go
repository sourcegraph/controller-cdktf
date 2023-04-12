package loadbalancerpool


type LoadBalancerPoolLoadShedding struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer_pool#default_percent LoadBalancerPool#default_percent}.
	DefaultPercent *float64 `field:"optional" json:"defaultPercent" yaml:"defaultPercent"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer_pool#default_policy LoadBalancerPool#default_policy}.
	DefaultPolicy *string `field:"optional" json:"defaultPolicy" yaml:"defaultPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer_pool#session_percent LoadBalancerPool#session_percent}.
	SessionPercent *float64 `field:"optional" json:"sessionPercent" yaml:"sessionPercent"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer_pool#session_policy LoadBalancerPool#session_policy}.
	SessionPolicy *string `field:"optional" json:"sessionPolicy" yaml:"sessionPolicy"`
}

