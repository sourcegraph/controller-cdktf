package googlecomputeinterconnectattachmentgroup


type GoogleComputeInterconnectAttachmentGroupAttachments struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_interconnect_attachment_group#name GoogleComputeInterconnectAttachmentGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_interconnect_attachment_group#attachment GoogleComputeInterconnectAttachmentGroup#attachment}.
	Attachment *string `field:"optional" json:"attachment" yaml:"attachment"`
}

