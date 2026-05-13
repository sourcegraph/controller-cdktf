package googlenetworksecurityullmirroringcollectorrule


type GoogleNetworkSecurityUllMirroringCollectorRuleMatch struct {
	// Direction of traffic to match.
	//
	// When unset, matches any direction.
	// Possible values:
	// INGRESS: Traffic inbound to the capture point.
	// EGRESS: Traffic outbound from the capture point.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_ull_mirroring_collector_rule#direction GoogleNetworkSecurityUllMirroringCollectorRule#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Destination IP ranges to match. When unset, matches any destination IP range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_ull_mirroring_collector_rule#dst_ip_ranges GoogleNetworkSecurityUllMirroringCollectorRule#dst_ip_ranges}
	DstIpRanges *[]*string `field:"optional" json:"dstIpRanges" yaml:"dstIpRanges"`
	// IP protocols to match. When unset, matches any IP protocol. Examples: "tcp", "udp", "icmp". If unset, matches any IP protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_ull_mirroring_collector_rule#ip_protocols GoogleNetworkSecurityUllMirroringCollectorRule#ip_protocols}
	IpProtocols *[]*string `field:"optional" json:"ipProtocols" yaml:"ipProtocols"`
	// Source IP ranges to match. When unset, matches any source IP range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_ull_mirroring_collector_rule#src_ip_ranges GoogleNetworkSecurityUllMirroringCollectorRule#src_ip_ranges}
	SrcIpRanges *[]*string `field:"optional" json:"srcIpRanges" yaml:"srcIpRanges"`
}

