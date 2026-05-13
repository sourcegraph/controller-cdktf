package googlecomputevpntunnel


type GoogleComputeVpnTunnelCipherSuite struct {
	// phase1 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_vpn_tunnel#phase1 GoogleComputeVpnTunnel#phase1}
	Phase1 *GoogleComputeVpnTunnelCipherSuitePhase1 `field:"optional" json:"phase1" yaml:"phase1"`
	// phase2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_vpn_tunnel#phase2 GoogleComputeVpnTunnel#phase2}
	Phase2 *GoogleComputeVpnTunnelCipherSuitePhase2 `field:"optional" json:"phase2" yaml:"phase2"`
}

