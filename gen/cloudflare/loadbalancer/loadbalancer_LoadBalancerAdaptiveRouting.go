package loadbalancer


type LoadBalancerAdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate pools, when no healthy alternate exists in the same pool, according to the failover order defined by traffic and origin steering.
	//
	// When set `false`, zero-downtime failover will only occur between origins within the same pool. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#failover_across_pools LoadBalancer#failover_across_pools}
	FailoverAcrossPools interface{} `field:"optional" json:"failoverAcrossPools" yaml:"failoverAcrossPools"`
}

