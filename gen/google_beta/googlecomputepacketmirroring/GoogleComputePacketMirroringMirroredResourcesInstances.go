package googlecomputepacketmirroring


type GoogleComputePacketMirroringMirroredResourcesInstances struct {
	// The URL of the instances where this rule should be active.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_packet_mirroring#url GoogleComputePacketMirroring#url}
	Url *string `field:"required" json:"url" yaml:"url"`
}

