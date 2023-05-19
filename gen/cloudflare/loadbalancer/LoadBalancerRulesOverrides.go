package loadbalancer


type LoadBalancerRulesOverrides struct {
	// adaptive_routing block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#adaptive_routing LoadBalancer#adaptive_routing}
	AdaptiveRouting interface{} `field:"optional" json:"adaptiveRouting" yaml:"adaptiveRouting"`
	// country_pools block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#country_pools LoadBalancer#country_pools}
	CountryPools interface{} `field:"optional" json:"countryPools" yaml:"countryPools"`
	// A list of pool IDs ordered by their failover priority. Used whenever [`pop_pools`](#pop_pools)/[`country_pools`](#country_pools)/[`region_pools`](#region_pools) are not defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#default_pools LoadBalancer#default_pools}
	DefaultPools *[]*string `field:"optional" json:"defaultPools" yaml:"defaultPools"`
	// The pool ID to use when all other pools are detected as unhealthy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#fallback_pool LoadBalancer#fallback_pool}
	FallbackPool *string `field:"optional" json:"fallbackPool" yaml:"fallbackPool"`
	// location_strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#location_strategy LoadBalancer#location_strategy}
	LocationStrategy interface{} `field:"optional" json:"locationStrategy" yaml:"locationStrategy"`
	// pop_pools block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#pop_pools LoadBalancer#pop_pools}
	PopPools interface{} `field:"optional" json:"popPools" yaml:"popPools"`
	// random_steering block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#random_steering LoadBalancer#random_steering}
	RandomSteering interface{} `field:"optional" json:"randomSteering" yaml:"randomSteering"`
	// region_pools block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#region_pools LoadBalancer#region_pools}
	RegionPools interface{} `field:"optional" json:"regionPools" yaml:"regionPools"`
	// Configure cookie attributes for session affinity cookie.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#session_affinity LoadBalancer#session_affinity}
	SessionAffinity *string `field:"optional" json:"sessionAffinity" yaml:"sessionAffinity"`
	// session_affinity_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#session_affinity_attributes LoadBalancer#session_affinity_attributes}
	SessionAffinityAttributes interface{} `field:"optional" json:"sessionAffinityAttributes" yaml:"sessionAffinityAttributes"`
	// Time, in seconds, until this load balancer's session affinity cookie expires after being created.
	//
	// This parameter is ignored unless a supported session affinity policy is set. The current default of `82800` (23 hours) will be used unless [`session_affinity_ttl`](#session_affinity_ttl) is explicitly set. Once the expiry time has been reached, subsequent requests may get sent to a different origin server. Valid values are between `1800` and `604800`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#session_affinity_ttl LoadBalancer#session_affinity_ttl}
	SessionAffinityTtl *float64 `field:"optional" json:"sessionAffinityTtl" yaml:"sessionAffinityTtl"`
	// The method the load balancer uses to determine the route to your origin.
	//
	// Value `off` uses [`default_pool_ids`](#default_pool_ids). Value `geo` uses [`pop_pools`](#pop_pools)/[`country_pools`](#country_pools)/[`region_pools`](#region_pools). For non-proxied requests, the [`country`](#country) for [`country_pools`](#country_pools) is determined by [`location_strategy`](#location_strategy). Value `random` selects a pool randomly. Value `dynamic_latency` uses round trip time to select the closest pool in [`default_pool_ids`](#default_pool_ids) (requires pool health checks). Value `proximity` uses the pools' latitude and longitude to select the closest pool using the Cloudflare PoP location for proxied requests or the location determined by [`location_strategy`](#location_strategy) for non-proxied requests. Value `""` maps to `geo` if you use [`pop_pools`](#pop_pools)/[`country_pools`](#country_pools)/[`region_pools`](#region_pools) otherwise `off`. Available values: `off`, `geo`, `dynamic_latency`, `random`, `proximity`, `""` Defaults to `""`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#steering_policy LoadBalancer#steering_policy}
	SteeringPolicy *string `field:"optional" json:"steeringPolicy" yaml:"steeringPolicy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load balancer.
	//
	// This cannot be set for proxied load balancers. Defaults to `30`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#ttl LoadBalancer#ttl}
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}

