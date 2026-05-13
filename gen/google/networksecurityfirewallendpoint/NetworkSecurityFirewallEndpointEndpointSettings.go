package networksecurityfirewallendpoint


type NetworkSecurityFirewallEndpointEndpointSettings struct {
	// Indicates whether Jumbo Frames are enabled for the firewall endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_security_firewall_endpoint#jumbo_frames_enabled NetworkSecurityFirewallEndpoint#jumbo_frames_enabled}
	JumboFramesEnabled interface{} `field:"optional" json:"jumboFramesEnabled" yaml:"jumboFramesEnabled"`
}

