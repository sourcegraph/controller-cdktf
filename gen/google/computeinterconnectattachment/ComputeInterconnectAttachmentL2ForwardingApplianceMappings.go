package computeinterconnectattachment


type ComputeInterconnectAttachmentL2ForwardingApplianceMappings struct {
	// The appliance IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_interconnect_attachment#appliance_ip_address ComputeInterconnectAttachment#appliance_ip_address}
	ApplianceIpAddress *string `field:"optional" json:"applianceIpAddress" yaml:"applianceIpAddress"`
	// inner_vlan_to_appliance_mappings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_interconnect_attachment#inner_vlan_to_appliance_mappings ComputeInterconnectAttachment#inner_vlan_to_appliance_mappings}
	InnerVlanToApplianceMappings interface{} `field:"optional" json:"innerVlanToApplianceMappings" yaml:"innerVlanToApplianceMappings"`
	// The name of this appliance mapping rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_interconnect_attachment#name ComputeInterconnectAttachment#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The VLAN tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_interconnect_attachment#vlan_id ComputeInterconnectAttachment#vlan_id}
	VlanId *string `field:"optional" json:"vlanId" yaml:"vlanId"`
}

