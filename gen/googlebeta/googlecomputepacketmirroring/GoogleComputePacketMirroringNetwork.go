package googlecomputepacketmirroring


type GoogleComputePacketMirroringNetwork struct {
	// The full self_link URL of the network where this rule is active.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_packet_mirroring#url GoogleComputePacketMirroring#url}
	Url *string `field:"required" json:"url" yaml:"url"`
}

