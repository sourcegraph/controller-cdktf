package loadbalancer


type LoadBalancerRegionPools struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#pool_ids LoadBalancer#pool_ids}.
	PoolIds *[]*string `field:"required" json:"poolIds" yaml:"poolIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#region LoadBalancer#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
}

