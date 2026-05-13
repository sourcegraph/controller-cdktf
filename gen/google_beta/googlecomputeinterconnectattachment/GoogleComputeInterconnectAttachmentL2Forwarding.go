package googlecomputeinterconnectattachment


type GoogleComputeInterconnectAttachmentL2Forwarding struct {
	// appliance_mappings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_interconnect_attachment#appliance_mappings GoogleComputeInterconnectAttachment#appliance_mappings}
	ApplianceMappings interface{} `field:"optional" json:"applianceMappings" yaml:"applianceMappings"`
	// The default appliance IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_interconnect_attachment#default_appliance_ip_address GoogleComputeInterconnectAttachment#default_appliance_ip_address}
	DefaultApplianceIpAddress *string `field:"optional" json:"defaultApplianceIpAddress" yaml:"defaultApplianceIpAddress"`
	// geneve_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_interconnect_attachment#geneve_header GoogleComputeInterconnectAttachment#geneve_header}
	GeneveHeader *GoogleComputeInterconnectAttachmentL2ForwardingGeneveHeader `field:"optional" json:"geneveHeader" yaml:"geneveHeader"`
	// URL of the network to which this attachment belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_interconnect_attachment#network GoogleComputeInterconnectAttachment#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// The tunnel endpoint IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_interconnect_attachment#tunnel_endpoint_ip_address GoogleComputeInterconnectAttachment#tunnel_endpoint_ip_address}
	TunnelEndpointIpAddress *string `field:"optional" json:"tunnelEndpointIpAddress" yaml:"tunnelEndpointIpAddress"`
}

