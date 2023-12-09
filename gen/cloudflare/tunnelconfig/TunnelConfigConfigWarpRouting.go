package tunnelconfig


type TunnelConfigConfigWarpRouting struct {
	// Whether WARP routing is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/tunnel_config#enabled TunnelConfigA#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

