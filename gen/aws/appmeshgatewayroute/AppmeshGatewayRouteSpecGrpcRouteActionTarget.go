package appmeshgatewayroute


type AppmeshGatewayRouteSpecGrpcRouteActionTarget struct {
	// virtual_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/appmesh_gateway_route#virtual_service AppmeshGatewayRoute#virtual_service}
	VirtualService *AppmeshGatewayRouteSpecGrpcRouteActionTargetVirtualService `field:"required" json:"virtualService" yaml:"virtualService"`
}

