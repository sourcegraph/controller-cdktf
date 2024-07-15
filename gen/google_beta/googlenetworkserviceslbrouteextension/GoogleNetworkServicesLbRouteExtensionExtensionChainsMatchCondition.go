package googlenetworkserviceslbrouteextension


type GoogleNetworkServicesLbRouteExtensionExtensionChainsMatchCondition struct {
	// A Common Expression Language (CEL) expression that is used to match requests for which the extension chain is executed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_lb_route_extension#cel_expression GoogleNetworkServicesLbRouteExtension#cel_expression}
	CelExpression *string `field:"required" json:"celExpression" yaml:"celExpression"`
}

