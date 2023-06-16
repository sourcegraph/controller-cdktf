package zonelockdown


type ZoneLockdownConfigurations struct {
	// The request property to target. Available values: `ip`, `ip_range`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_lockdown#target ZoneLockdown#target}
	Target *string `field:"required" json:"target" yaml:"target"`
	// The value to target.
	//
	// Depends on target's type. IP addresses should just be standard IPv4/IPv6 notation i.e. `192.0.2.1` or `2001:db8::/32` and IP ranges in CIDR format i.e. `192.0.2.0/24`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_lockdown#value ZoneLockdown#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

