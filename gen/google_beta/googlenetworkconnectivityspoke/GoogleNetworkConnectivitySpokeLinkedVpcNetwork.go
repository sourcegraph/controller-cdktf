package googlenetworkconnectivityspoke


type GoogleNetworkConnectivitySpokeLinkedVpcNetwork struct {
	// The URI of the VPC network resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_network_connectivity_spoke#uri GoogleNetworkConnectivitySpoke#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// IP ranges encompassing the subnets to be excluded from peering.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_network_connectivity_spoke#exclude_export_ranges GoogleNetworkConnectivitySpoke#exclude_export_ranges}
	ExcludeExportRanges *[]*string `field:"optional" json:"excludeExportRanges" yaml:"excludeExportRanges"`
	// IP ranges allowed to be included from peering.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_network_connectivity_spoke#include_export_ranges GoogleNetworkConnectivitySpoke#include_export_ranges}
	IncludeExportRanges *[]*string `field:"optional" json:"includeExportRanges" yaml:"includeExportRanges"`
}

