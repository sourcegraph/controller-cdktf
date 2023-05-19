package tunnelconfig


type TunnelConfigConfigOriginRequestIpRules struct {
	// Whether to allow the IP prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/tunnel_config#allow TunnelConfigA#allow}
	Allow interface{} `field:"optional" json:"allow" yaml:"allow"`
	// Ports to use within the IP rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/tunnel_config#ports TunnelConfigA#ports}
	Ports *[]*float64 `field:"optional" json:"ports" yaml:"ports"`
	// IP rule prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/tunnel_config#prefix TunnelConfigA#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

