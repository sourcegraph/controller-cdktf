package googlenetworkserviceslbedgeextension


type GoogleNetworkServicesLbEdgeExtensionExtensionChainsMatchCondition struct {
	// A Common Expression Language (CEL) expression that is used to match requests for which the extension chain is executed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_lb_edge_extension#cel_expression GoogleNetworkServicesLbEdgeExtension#cel_expression}
	CelExpression *string `field:"required" json:"celExpression" yaml:"celExpression"`
}

