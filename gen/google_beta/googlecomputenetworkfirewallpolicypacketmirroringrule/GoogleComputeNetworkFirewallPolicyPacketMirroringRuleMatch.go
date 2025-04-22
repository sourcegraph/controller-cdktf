package googlecomputenetworkfirewallpolicypacketmirroringrule


type GoogleComputeNetworkFirewallPolicyPacketMirroringRuleMatch struct {
	// layer4_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_network_firewall_policy_packet_mirroring_rule#layer4_configs GoogleComputeNetworkFirewallPolicyPacketMirroringRule#layer4_configs}
	Layer4Configs interface{} `field:"required" json:"layer4Configs" yaml:"layer4Configs"`
	// CIDR IP address range. Maximum number of destination CIDR IP ranges allowed is 5000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_network_firewall_policy_packet_mirroring_rule#dest_ip_ranges GoogleComputeNetworkFirewallPolicyPacketMirroringRule#dest_ip_ranges}
	DestIpRanges *[]*string `field:"optional" json:"destIpRanges" yaml:"destIpRanges"`
	// CIDR IP address range. Maximum number of source CIDR IP ranges allowed is 5000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_network_firewall_policy_packet_mirroring_rule#src_ip_ranges GoogleComputeNetworkFirewallPolicyPacketMirroringRule#src_ip_ranges}
	SrcIpRanges *[]*string `field:"optional" json:"srcIpRanges" yaml:"srcIpRanges"`
}

