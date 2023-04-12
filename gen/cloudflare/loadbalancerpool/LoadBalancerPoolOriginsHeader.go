package loadbalancerpool


type LoadBalancerPoolOriginsHeader struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer_pool#header LoadBalancerPool#header}.
	Header *string `field:"required" json:"header" yaml:"header"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer_pool#values LoadBalancerPool#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

