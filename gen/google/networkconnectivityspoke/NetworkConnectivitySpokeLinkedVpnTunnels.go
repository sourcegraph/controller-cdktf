package networkconnectivityspoke


type NetworkConnectivitySpokeLinkedVpnTunnels struct {
	// A value that controls whether site-to-site data transfer is enabled for these resources.
	//
	// Note that data transfer is available only in supported locations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_connectivity_spoke#site_to_site_data_transfer NetworkConnectivitySpoke#site_to_site_data_transfer}
	SiteToSiteDataTransfer interface{} `field:"required" json:"siteToSiteDataTransfer" yaml:"siteToSiteDataTransfer"`
	// The URIs of linked VPN tunnel resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_connectivity_spoke#uris NetworkConnectivitySpoke#uris}
	Uris *[]*string `field:"required" json:"uris" yaml:"uris"`
	// Dynamic routes overlapped/encompassed by exclude export ranges are excluded during export to hub.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_connectivity_spoke#exclude_export_ranges NetworkConnectivitySpoke#exclude_export_ranges}
	ExcludeExportRanges *[]*string `field:"optional" json:"excludeExportRanges" yaml:"excludeExportRanges"`
	// Hub routes overlapped/encompassed by exclude import ranges are excluded during import from hub.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_connectivity_spoke#exclude_import_ranges NetworkConnectivitySpoke#exclude_import_ranges}
	ExcludeImportRanges *[]*string `field:"optional" json:"excludeImportRanges" yaml:"excludeImportRanges"`
	// Dynamic routes fully encompassed by include export ranges are included during export to hub.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_connectivity_spoke#include_export_ranges NetworkConnectivitySpoke#include_export_ranges}
	IncludeExportRanges *[]*string `field:"optional" json:"includeExportRanges" yaml:"includeExportRanges"`
	// Hub routes fully encompassed by include import ranges are included during import from hub.
	//
	// "ALL_IPV4_RANGES" or IPv4 CIDR ranges are allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_connectivity_spoke#include_import_ranges NetworkConnectivitySpoke#include_import_ranges}
	IncludeImportRanges *[]*string `field:"optional" json:"includeImportRanges" yaml:"includeImportRanges"`
}

