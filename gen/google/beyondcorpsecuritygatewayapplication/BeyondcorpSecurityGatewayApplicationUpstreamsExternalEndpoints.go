package beyondcorpsecuritygatewayapplication


type BeyondcorpSecurityGatewayApplicationUpstreamsExternalEndpoints struct {
	// Hostname of the endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/beyondcorp_security_gateway_application#hostname BeyondcorpSecurityGatewayApplication#hostname}
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// Port of the endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/beyondcorp_security_gateway_application#port BeyondcorpSecurityGatewayApplication#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

