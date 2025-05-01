package googlenetworkconnectivityspoke


type GoogleNetworkConnectivitySpokeLinkedProducerVpcNetwork struct {
	// The URI of the Service Consumer VPC that the Producer VPC is peered with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_network_connectivity_spoke#network GoogleNetworkConnectivitySpoke#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// The name of the VPC peering between the Service Consumer VPC and the Producer VPC (defined in the Tenant project) which is added to the NCC hub.
	//
	// This peering must be in ACTIVE state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_network_connectivity_spoke#peering GoogleNetworkConnectivitySpoke#peering}
	Peering *string `field:"required" json:"peering" yaml:"peering"`
	// IP ranges encompassing the subnets to be excluded from peering.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_network_connectivity_spoke#exclude_export_ranges GoogleNetworkConnectivitySpoke#exclude_export_ranges}
	ExcludeExportRanges *[]*string `field:"optional" json:"excludeExportRanges" yaml:"excludeExportRanges"`
	// IP ranges allowed to be included from peering.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_network_connectivity_spoke#include_export_ranges GoogleNetworkConnectivitySpoke#include_export_ranges}
	IncludeExportRanges *[]*string `field:"optional" json:"includeExportRanges" yaml:"includeExportRanges"`
}

