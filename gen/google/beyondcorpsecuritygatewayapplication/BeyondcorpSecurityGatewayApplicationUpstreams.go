package beyondcorpsecuritygatewayapplication


type BeyondcorpSecurityGatewayApplicationUpstreams struct {
	// egress_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/beyondcorp_security_gateway_application#egress_policy BeyondcorpSecurityGatewayApplication#egress_policy}
	EgressPolicy *BeyondcorpSecurityGatewayApplicationUpstreamsEgressPolicy `field:"optional" json:"egressPolicy" yaml:"egressPolicy"`
	// external block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/beyondcorp_security_gateway_application#external BeyondcorpSecurityGatewayApplication#external}
	External *BeyondcorpSecurityGatewayApplicationUpstreamsExternal `field:"optional" json:"external" yaml:"external"`
	// network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/beyondcorp_security_gateway_application#network BeyondcorpSecurityGatewayApplication#network}
	Network *BeyondcorpSecurityGatewayApplicationUpstreamsNetwork `field:"optional" json:"network" yaml:"network"`
	// proxy_protocol block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/beyondcorp_security_gateway_application#proxy_protocol BeyondcorpSecurityGatewayApplication#proxy_protocol}
	ProxyProtocol *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocol `field:"optional" json:"proxyProtocol" yaml:"proxyProtocol"`
}

