package googlecomputeregionnetworkfirewallpolicywithrules


type GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatch struct {
	// layer4_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_region_network_firewall_policy_with_rules#layer4_config GoogleComputeRegionNetworkFirewallPolicyWithRules#layer4_config}
	Layer4Config interface{} `field:"required" json:"layer4Config" yaml:"layer4Config"`
	// Address groups which should be matched against the traffic destination. Maximum number of destination address groups is 10.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_region_network_firewall_policy_with_rules#dest_address_groups GoogleComputeRegionNetworkFirewallPolicyWithRules#dest_address_groups}
	DestAddressGroups *[]*string `field:"optional" json:"destAddressGroups" yaml:"destAddressGroups"`
	// Fully Qualified Domain Name (FQDN) which should be matched against traffic destination. Maximum number of destination fqdn allowed is 100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_region_network_firewall_policy_with_rules#dest_fqdns GoogleComputeRegionNetworkFirewallPolicyWithRules#dest_fqdns}
	DestFqdns *[]*string `field:"optional" json:"destFqdns" yaml:"destFqdns"`
	// Destination IP address range in CIDR format. Required for EGRESS rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_region_network_firewall_policy_with_rules#dest_ip_ranges GoogleComputeRegionNetworkFirewallPolicyWithRules#dest_ip_ranges}
	DestIpRanges *[]*string `field:"optional" json:"destIpRanges" yaml:"destIpRanges"`
	// Network scope of the traffic destination. Possible values: ["INTERNET", "INTRA_VPC", "NON_INTERNET", "VPC_NETWORKS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_region_network_firewall_policy_with_rules#dest_network_scope GoogleComputeRegionNetworkFirewallPolicyWithRules#dest_network_scope}
	DestNetworkScope *string `field:"optional" json:"destNetworkScope" yaml:"destNetworkScope"`
	// Region codes whose IP addresses will be used to match for destination of traffic.
	//
	// Should be specified as 2 letter country code defined as per
	// ISO 3166 alpha-2 country codes. ex."US"
	// Maximum number of destination region codes allowed is 5000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_region_network_firewall_policy_with_rules#dest_region_codes GoogleComputeRegionNetworkFirewallPolicyWithRules#dest_region_codes}
	DestRegionCodes *[]*string `field:"optional" json:"destRegionCodes" yaml:"destRegionCodes"`
	// Names of Network Threat Intelligence lists. The IPs in these lists will be matched against traffic destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_region_network_firewall_policy_with_rules#dest_threat_intelligences GoogleComputeRegionNetworkFirewallPolicyWithRules#dest_threat_intelligences}
	DestThreatIntelligences *[]*string `field:"optional" json:"destThreatIntelligences" yaml:"destThreatIntelligences"`
	// Address groups which should be matched against the traffic source. Maximum number of source address groups is 10.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_region_network_firewall_policy_with_rules#src_address_groups GoogleComputeRegionNetworkFirewallPolicyWithRules#src_address_groups}
	SrcAddressGroups *[]*string `field:"optional" json:"srcAddressGroups" yaml:"srcAddressGroups"`
	// Fully Qualified Domain Name (FQDN) which should be matched against traffic source. Maximum number of source fqdn allowed is 100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_region_network_firewall_policy_with_rules#src_fqdns GoogleComputeRegionNetworkFirewallPolicyWithRules#src_fqdns}
	SrcFqdns *[]*string `field:"optional" json:"srcFqdns" yaml:"srcFqdns"`
	// Source IP address range in CIDR format. Required for INGRESS rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_region_network_firewall_policy_with_rules#src_ip_ranges GoogleComputeRegionNetworkFirewallPolicyWithRules#src_ip_ranges}
	SrcIpRanges *[]*string `field:"optional" json:"srcIpRanges" yaml:"srcIpRanges"`
	// Networks of the traffic source. It can be either a full or partial url.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_region_network_firewall_policy_with_rules#src_networks GoogleComputeRegionNetworkFirewallPolicyWithRules#src_networks}
	SrcNetworks *[]*string `field:"optional" json:"srcNetworks" yaml:"srcNetworks"`
	// Network scope of the traffic source. Possible values: ["INTERNET", "INTRA_VPC", "NON_INTERNET", "VPC_NETWORKS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_region_network_firewall_policy_with_rules#src_network_scope GoogleComputeRegionNetworkFirewallPolicyWithRules#src_network_scope}
	SrcNetworkScope *string `field:"optional" json:"srcNetworkScope" yaml:"srcNetworkScope"`
	// Region codes whose IP addresses will be used to match for source of traffic.
	//
	// Should be specified as 2 letter country code defined as per
	// ISO 3166 alpha-2 country codes. ex."US"
	// Maximum number of source region codes allowed is 5000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_region_network_firewall_policy_with_rules#src_region_codes GoogleComputeRegionNetworkFirewallPolicyWithRules#src_region_codes}
	SrcRegionCodes *[]*string `field:"optional" json:"srcRegionCodes" yaml:"srcRegionCodes"`
	// src_secure_tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_region_network_firewall_policy_with_rules#src_secure_tag GoogleComputeRegionNetworkFirewallPolicyWithRules#src_secure_tag}
	SrcSecureTag interface{} `field:"optional" json:"srcSecureTag" yaml:"srcSecureTag"`
	// Names of Network Threat Intelligence lists. The IPs in these lists will be matched against traffic source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_region_network_firewall_policy_with_rules#src_threat_intelligences GoogleComputeRegionNetworkFirewallPolicyWithRules#src_threat_intelligences}
	SrcThreatIntelligences *[]*string `field:"optional" json:"srcThreatIntelligences" yaml:"srcThreatIntelligences"`
}

