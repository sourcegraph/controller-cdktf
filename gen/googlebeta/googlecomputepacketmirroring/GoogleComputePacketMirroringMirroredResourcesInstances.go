package googlecomputepacketmirroring


type GoogleComputePacketMirroringMirroredResourcesInstances struct {
	// The URL of the instances where this rule should be active.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_packet_mirroring#url GoogleComputePacketMirroring#url}
	Url *string `field:"required" json:"url" yaml:"url"`
}

