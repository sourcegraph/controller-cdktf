package teamsrule


type TeamsRuleRuleSettingsEgress struct {
	// The IPv4 address to be used for egress.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_rule#ipv4 TeamsRule#ipv4}
	Ipv4 *string `field:"required" json:"ipv4" yaml:"ipv4"`
	// The IPv6 range to be used for egress.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_rule#ipv6 TeamsRule#ipv6}
	Ipv6 *string `field:"required" json:"ipv6" yaml:"ipv6"`
	// The IPv4 address to be used for egress in the event of an error egressing with the primary IPv4.
	//
	// Can be '0.0.0.0' to indicate local egreass via Warp IPs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_rule#ipv4_fallback TeamsRule#ipv4_fallback}
	Ipv4Fallback *string `field:"optional" json:"ipv4Fallback" yaml:"ipv4Fallback"`
}

