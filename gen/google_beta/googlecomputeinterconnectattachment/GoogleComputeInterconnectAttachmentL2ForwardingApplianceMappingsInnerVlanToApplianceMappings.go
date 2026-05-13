package googlecomputeinterconnectattachment


type GoogleComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappings struct {
	// The inner appliance IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_interconnect_attachment#inner_appliance_ip_address GoogleComputeInterconnectAttachment#inner_appliance_ip_address}
	InnerApplianceIpAddress *string `field:"optional" json:"innerApplianceIpAddress" yaml:"innerApplianceIpAddress"`
	// List of inner VLAN tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_interconnect_attachment#inner_vlan_tags GoogleComputeInterconnectAttachment#inner_vlan_tags}
	InnerVlanTags *[]*string `field:"optional" json:"innerVlanTags" yaml:"innerVlanTags"`
}

