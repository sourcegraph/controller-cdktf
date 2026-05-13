package computeinterconnectattachment


type ComputeInterconnectAttachmentL2ForwardingGeneveHeader struct {
	// VNI is a 24-bit unique virtual network identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_interconnect_attachment#vni ComputeInterconnectAttachment#vni}
	Vni *float64 `field:"optional" json:"vni" yaml:"vni"`
}

