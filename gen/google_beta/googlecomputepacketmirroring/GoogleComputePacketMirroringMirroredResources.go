package googlecomputepacketmirroring


type GoogleComputePacketMirroringMirroredResources struct {
	// instances block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_packet_mirroring#instances GoogleComputePacketMirroring#instances}
	Instances interface{} `field:"optional" json:"instances" yaml:"instances"`
	// subnetworks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_packet_mirroring#subnetworks GoogleComputePacketMirroring#subnetworks}
	Subnetworks interface{} `field:"optional" json:"subnetworks" yaml:"subnetworks"`
	// All instances with these tags will be mirrored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_packet_mirroring#tags GoogleComputePacketMirroring#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

