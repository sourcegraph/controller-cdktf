package computevpntunnel


type ComputeVpnTunnelCipherSuitePhase1 struct {
	// Diffie-Hellman groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_vpn_tunnel#dh ComputeVpnTunnel#dh}
	Dh *[]*string `field:"optional" json:"dh" yaml:"dh"`
	// Encryption algorithms.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_vpn_tunnel#encryption ComputeVpnTunnel#encryption}
	Encryption *[]*string `field:"optional" json:"encryption" yaml:"encryption"`
	// Integrity algorithms.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_vpn_tunnel#integrity ComputeVpnTunnel#integrity}
	Integrity *[]*string `field:"optional" json:"integrity" yaml:"integrity"`
	// Pseudo-random functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_vpn_tunnel#prf ComputeVpnTunnel#prf}
	Prf *[]*string `field:"optional" json:"prf" yaml:"prf"`
}

