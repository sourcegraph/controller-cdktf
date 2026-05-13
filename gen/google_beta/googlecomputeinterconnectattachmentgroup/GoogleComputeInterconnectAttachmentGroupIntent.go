package googlecomputeinterconnectattachmentgroup


type GoogleComputeInterconnectAttachmentGroupIntent struct {
	// Which SLA the user intends this group to support. Possible values: ["PRODUCTION_NON_CRITICAL", "PRODUCTION_CRITICAL", "NO_SLA", "AVAILABILITY_SLA_UNSPECIFIED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_interconnect_attachment_group#availability_sla GoogleComputeInterconnectAttachmentGroup#availability_sla}
	AvailabilitySla *string `field:"optional" json:"availabilitySla" yaml:"availabilitySla"`
}

