package googlenetworkserviceslbtrafficextension


type GoogleNetworkServicesLbTrafficExtensionExtensionChainsMatchCondition struct {
	// A Common Expression Language (CEL) expression that is used to match requests for which the extension chain is executed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_lb_traffic_extension#cel_expression GoogleNetworkServicesLbTrafficExtension#cel_expression}
	CelExpression *string `field:"required" json:"celExpression" yaml:"celExpression"`
}

