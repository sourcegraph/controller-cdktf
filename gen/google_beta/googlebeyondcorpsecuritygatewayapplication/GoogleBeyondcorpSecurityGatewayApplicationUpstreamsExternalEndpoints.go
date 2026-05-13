package googlebeyondcorpsecuritygatewayapplication


type GoogleBeyondcorpSecurityGatewayApplicationUpstreamsExternalEndpoints struct {
	// Hostname of the endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_beyondcorp_security_gateway_application#hostname GoogleBeyondcorpSecurityGatewayApplication#hostname}
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// Port of the endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_beyondcorp_security_gateway_application#port GoogleBeyondcorpSecurityGatewayApplication#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

