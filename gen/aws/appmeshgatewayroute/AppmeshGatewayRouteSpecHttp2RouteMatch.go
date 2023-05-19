package appmeshgatewayroute


type AppmeshGatewayRouteSpecHttp2RouteMatch struct {
	// hostname block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/appmesh_gateway_route#hostname AppmeshGatewayRoute#hostname}
	Hostname *AppmeshGatewayRouteSpecHttp2RouteMatchHostname `field:"optional" json:"hostname" yaml:"hostname"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/appmesh_gateway_route#port AppmeshGatewayRoute#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/appmesh_gateway_route#prefix AppmeshGatewayRoute#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

