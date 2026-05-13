package beyondcorpsecuritygatewayapplication


type BeyondcorpSecurityGatewayApplicationEndpointMatchers struct {
	// Required. Hostname of the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/beyondcorp_security_gateway_application#hostname BeyondcorpSecurityGatewayApplication#hostname}
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// Optional. Ports of the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/beyondcorp_security_gateway_application#ports BeyondcorpSecurityGatewayApplication#ports}
	Ports *[]*float64 `field:"required" json:"ports" yaml:"ports"`
}

