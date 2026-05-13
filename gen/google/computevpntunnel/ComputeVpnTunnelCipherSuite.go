package computevpntunnel


type ComputeVpnTunnelCipherSuite struct {
	// phase1 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_vpn_tunnel#phase1 ComputeVpnTunnel#phase1}
	Phase1 *ComputeVpnTunnelCipherSuitePhase1 `field:"optional" json:"phase1" yaml:"phase1"`
	// phase2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_vpn_tunnel#phase2 ComputeVpnTunnel#phase2}
	Phase2 *ComputeVpnTunnelCipherSuitePhase2 `field:"optional" json:"phase2" yaml:"phase2"`
}

