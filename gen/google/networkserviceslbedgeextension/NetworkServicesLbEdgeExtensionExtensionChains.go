package networkserviceslbedgeextension


type NetworkServicesLbEdgeExtensionExtensionChains struct {
	// extensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_lb_edge_extension#extensions NetworkServicesLbEdgeExtension#extensions}
	Extensions interface{} `field:"required" json:"extensions" yaml:"extensions"`
	// match_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_lb_edge_extension#match_condition NetworkServicesLbEdgeExtension#match_condition}
	MatchCondition *NetworkServicesLbEdgeExtensionExtensionChainsMatchCondition `field:"required" json:"matchCondition" yaml:"matchCondition"`
	// The name for this extension chain.
	//
	// The name is logged as part of the HTTP request logs.
	// The name must conform with RFC-1034, is restricted to lower-cased letters, numbers and hyphens,
	// and can have a maximum length of 63 characters. Additionally, the first character must be a letter
	// and the last character must be a letter or a number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_lb_edge_extension#name NetworkServicesLbEdgeExtension#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

