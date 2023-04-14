package tunnelconfig


type TunnelConfigConfigWarpRouting struct {
	// Whether WARP routing is enabled.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/tunnel_config#enabled TunnelConfigA#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

