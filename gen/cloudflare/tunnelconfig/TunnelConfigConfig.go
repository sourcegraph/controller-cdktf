package tunnelconfig


type TunnelConfigConfig struct {
	// ingress_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/tunnel_config#ingress_rule TunnelConfigA#ingress_rule}
	IngressRule interface{} `field:"required" json:"ingressRule" yaml:"ingressRule"`
	// origin_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/tunnel_config#origin_request TunnelConfigA#origin_request}
	OriginRequest *TunnelConfigConfigOriginRequest `field:"optional" json:"originRequest" yaml:"originRequest"`
	// warp_routing block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/tunnel_config#warp_routing TunnelConfigA#warp_routing}
	WarpRouting *TunnelConfigConfigWarpRouting `field:"optional" json:"warpRouting" yaml:"warpRouting"`
}

