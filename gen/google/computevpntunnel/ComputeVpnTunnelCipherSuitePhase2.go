package computevpntunnel


type ComputeVpnTunnelCipherSuitePhase2 struct {
	// Encryption algorithms.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_vpn_tunnel#encryption ComputeVpnTunnel#encryption}
	Encryption *[]*string `field:"optional" json:"encryption" yaml:"encryption"`
	// Integrity algorithms.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_vpn_tunnel#integrity ComputeVpnTunnel#integrity}
	Integrity *[]*string `field:"optional" json:"integrity" yaml:"integrity"`
	// Perfect forward secrecy groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_vpn_tunnel#pfs ComputeVpnTunnel#pfs}
	Pfs *[]*string `field:"optional" json:"pfs" yaml:"pfs"`
}

