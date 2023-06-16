package splittunnel


type SplitTunnelTunnels struct {
	// The address for the tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/split_tunnel#address SplitTunnel#address}
	Address *string `field:"optional" json:"address" yaml:"address"`
	// A description for the tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/split_tunnel#description SplitTunnel#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The domain name for the tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/split_tunnel#host SplitTunnel#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
}

