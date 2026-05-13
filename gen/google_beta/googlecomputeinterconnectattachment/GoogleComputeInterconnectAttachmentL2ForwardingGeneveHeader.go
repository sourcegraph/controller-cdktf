package googlecomputeinterconnectattachment


type GoogleComputeInterconnectAttachmentL2ForwardingGeneveHeader struct {
	// VNI is a 24-bit unique virtual network identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_interconnect_attachment#vni GoogleComputeInterconnectAttachment#vni}
	Vni *float64 `field:"optional" json:"vni" yaml:"vni"`
}

