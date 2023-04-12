package loadbalancer


type LoadBalancerPopPools struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#pool_ids LoadBalancer#pool_ids}.
	PoolIds *[]*string `field:"required" json:"poolIds" yaml:"poolIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#pop LoadBalancer#pop}.
	Pop *string `field:"required" json:"pop" yaml:"pop"`
}

