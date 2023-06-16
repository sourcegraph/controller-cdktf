package fallbackdomain


type FallbackDomainDomains struct {
	// A description of the fallback domain, displayed in the client UI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/fallback_domain#description FallbackDomain#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A list of IP addresses to handle domain resolution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/fallback_domain#dns_server FallbackDomain#dns_server}
	DnsServer *[]*string `field:"optional" json:"dnsServer" yaml:"dnsServer"`
	// The domain suffix to match when resolving locally.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/fallback_domain#suffix FallbackDomain#suffix}
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
}

