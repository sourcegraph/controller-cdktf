package googlecomputevpntunnel


type GoogleComputeVpnTunnelCipherSuitePhase1 struct {
	// Diffie-Hellman groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_vpn_tunnel#dh GoogleComputeVpnTunnel#dh}
	Dh *[]*string `field:"optional" json:"dh" yaml:"dh"`
	// Encryption algorithms.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_vpn_tunnel#encryption GoogleComputeVpnTunnel#encryption}
	Encryption *[]*string `field:"optional" json:"encryption" yaml:"encryption"`
	// Integrity algorithms.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_vpn_tunnel#integrity GoogleComputeVpnTunnel#integrity}
	Integrity *[]*string `field:"optional" json:"integrity" yaml:"integrity"`
	// Pseudo-random functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_vpn_tunnel#prf GoogleComputeVpnTunnel#prf}
	Prf *[]*string `field:"optional" json:"prf" yaml:"prf"`
}

