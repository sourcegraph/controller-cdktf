package loadbalancer


type LoadBalancerPopPools struct {
	// A list of pool IDs in failover priority to use for traffic reaching the given PoP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#pool_ids LoadBalancer#pool_ids}
	PoolIds *[]*string `field:"required" json:"poolIds" yaml:"poolIds"`
	// A 3-letter code for the Point-of-Presence.
	//
	// Allowed values can be found in the list of datacenters on the [status page](https://www.cloudflarestatus.com/). Multiple entries should not be specified with the same PoP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#pop LoadBalancer#pop}
	Pop *string `field:"required" json:"pop" yaml:"pop"`
}

