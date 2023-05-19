package loadbalancer


type LoadBalancerRulesOverridesSessionAffinityAttributes struct {
	// Configures the SameSite attribute on session affinity cookie.
	//
	// Value `Auto` will be translated to `Lax` or `None` depending if Always Use HTTPS is enabled. Note: when using value `None`, then you can not set [`secure="Never"`](#secure). Available values: `Auto`, `Lax`, `None`, `Strict`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#samesite LoadBalancer#samesite}
	Samesite *string `field:"optional" json:"samesite" yaml:"samesite"`
	// Configures the Secure attribute on session affinity cookie.
	//
	// Value `Always` indicates the Secure attribute will be set in the Set-Cookie header, `Never` indicates the Secure attribute will not be set, and `Auto` will set the Secure attribute depending if Always Use HTTPS is enabled. Available values: `Auto`, `Always`, `Never`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#secure LoadBalancer#secure}
	Secure *string `field:"optional" json:"secure" yaml:"secure"`
	// Configures the zero-downtime failover between origins within a pool when session affinity is enabled.
	//
	// Value `none` means no failover takes place for sessions pinned to the origin. Value `temporary` means traffic will be sent to another other healthy origin until the originally pinned origin is available; note that this can potentially result in heavy origin flapping. Value `sticky` means the session affinity cookie is updated and subsequent requests are sent to the new origin. This feature is currently incompatible with Argo, Tiered Cache, and Bandwidth Alliance. Available values: `none`, `temporary`, `sticky`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#zero_downtime_failover LoadBalancer#zero_downtime_failover}
	ZeroDowntimeFailover *string `field:"optional" json:"zeroDowntimeFailover" yaml:"zeroDowntimeFailover"`
}

