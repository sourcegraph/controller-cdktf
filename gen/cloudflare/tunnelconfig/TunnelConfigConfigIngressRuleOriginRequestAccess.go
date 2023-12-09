package tunnelconfig


type TunnelConfigConfigIngressRuleOriginRequestAccess struct {
	// Audience tags of the access rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/tunnel_config#aud_tag TunnelConfigA#aud_tag}
	AudTag *[]*string `field:"optional" json:"audTag" yaml:"audTag"`
	// Whether the access rule is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/tunnel_config#required TunnelConfigA#required}
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// Name of the team to which the access rule applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/tunnel_config#team_name TunnelConfigA#team_name}
	TeamName *string `field:"optional" json:"teamName" yaml:"teamName"`
}

