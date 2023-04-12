package loadbalancerpool


type LoadBalancerPoolOriginSteering struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer_pool#policy LoadBalancerPool#policy}.
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
}

