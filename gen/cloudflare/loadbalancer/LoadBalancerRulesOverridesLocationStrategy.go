package loadbalancer


type LoadBalancerRulesOverridesLocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist in the request, or its GeoIP lookup is unsuccessful.
	//
	// Value `pop` will use the Cloudflare PoP location. Value `resolver_ip` will use the DNS resolver GeoIP location. If the GeoIP lookup is unsuccessful, it will use the Cloudflare PoP location. Available values: `pop`, `resolver_ip`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#mode LoadBalancer#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the authoritative location.
	//
	// Value `always` will always prefer ECS, `never` will never prefer ECS, `proximity` will prefer ECS only when [`steering_policy="proximity"`](#steering_policy), and `geo` will prefer ECS only when [`steering_policy="geo"`](#steering_policy). Available values: `always`, `never`, `proximity`, `geo`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#prefer_ecs LoadBalancer#prefer_ecs}
	PreferEcs *string `field:"optional" json:"preferEcs" yaml:"preferEcs"`
}

