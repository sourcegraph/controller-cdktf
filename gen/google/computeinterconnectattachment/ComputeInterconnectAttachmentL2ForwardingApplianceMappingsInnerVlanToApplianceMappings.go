package computeinterconnectattachment


type ComputeInterconnectAttachmentL2ForwardingApplianceMappingsInnerVlanToApplianceMappings struct {
	// The inner appliance IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_interconnect_attachment#inner_appliance_ip_address ComputeInterconnectAttachment#inner_appliance_ip_address}
	InnerApplianceIpAddress *string `field:"optional" json:"innerApplianceIpAddress" yaml:"innerApplianceIpAddress"`
	// List of inner VLAN tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_interconnect_attachment#inner_vlan_tags ComputeInterconnectAttachment#inner_vlan_tags}
	InnerVlanTags *[]*string `field:"optional" json:"innerVlanTags" yaml:"innerVlanTags"`
}

