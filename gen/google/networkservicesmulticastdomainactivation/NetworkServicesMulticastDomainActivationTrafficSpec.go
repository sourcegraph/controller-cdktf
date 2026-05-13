package networkservicesmulticastdomainactivation


type NetworkServicesMulticastDomainActivationTrafficSpec struct {
	// Aggregated egress Packet-Per-Second for all multicast groups in the domain in this zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_multicast_domain_activation#aggr_egress_pps NetworkServicesMulticastDomainActivation#aggr_egress_pps}
	AggrEgressPps *string `field:"optional" json:"aggrEgressPps" yaml:"aggrEgressPps"`
	// Aggregated ingress Packet-Per-Second for all multicast groups in the domain in this zone. Default to (aggregated_egress_pps / max_per_group_subscribers) * 2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_multicast_domain_activation#aggr_ingress_pps NetworkServicesMulticastDomainActivation#aggr_ingress_pps}
	AggrIngressPps *string `field:"optional" json:"aggrIngressPps" yaml:"aggrIngressPps"`
	// Average packet size (Default to 512 bytes).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_multicast_domain_activation#avg_packet_size NetworkServicesMulticastDomainActivation#avg_packet_size}
	AvgPacketSize *float64 `field:"optional" json:"avgPacketSize" yaml:"avgPacketSize"`
	// Maximum ingress Packet-Per-Second for a single multicast group in this zone. Default to aggregated_ingress_pps / 2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_multicast_domain_activation#max_per_group_ingress_pps NetworkServicesMulticastDomainActivation#max_per_group_ingress_pps}
	MaxPerGroupIngressPps *string `field:"optional" json:"maxPerGroupIngressPps" yaml:"maxPerGroupIngressPps"`
	// Maximum number of subscribers for a single multicast group in this zone. Default to max(50, aggregated_egress_pps / aggregated_ingress_pps).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_multicast_domain_activation#max_per_group_subscribers NetworkServicesMulticastDomainActivation#max_per_group_subscribers}
	MaxPerGroupSubscribers *string `field:"optional" json:"maxPerGroupSubscribers" yaml:"maxPerGroupSubscribers"`
}

