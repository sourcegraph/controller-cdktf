package googlebeyondcorpsecuritygatewayapplication


type GoogleBeyondcorpSecurityGatewayApplicationUpstreams struct {
	// egress_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_beyondcorp_security_gateway_application#egress_policy GoogleBeyondcorpSecurityGatewayApplication#egress_policy}
	EgressPolicy *GoogleBeyondcorpSecurityGatewayApplicationUpstreamsEgressPolicy `field:"optional" json:"egressPolicy" yaml:"egressPolicy"`
	// external block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_beyondcorp_security_gateway_application#external GoogleBeyondcorpSecurityGatewayApplication#external}
	External *GoogleBeyondcorpSecurityGatewayApplicationUpstreamsExternal `field:"optional" json:"external" yaml:"external"`
	// network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_beyondcorp_security_gateway_application#network GoogleBeyondcorpSecurityGatewayApplication#network}
	Network *GoogleBeyondcorpSecurityGatewayApplicationUpstreamsNetwork `field:"optional" json:"network" yaml:"network"`
	// proxy_protocol block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_beyondcorp_security_gateway_application#proxy_protocol GoogleBeyondcorpSecurityGatewayApplication#proxy_protocol}
	ProxyProtocol *GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocol `field:"optional" json:"proxyProtocol" yaml:"proxyProtocol"`
}

