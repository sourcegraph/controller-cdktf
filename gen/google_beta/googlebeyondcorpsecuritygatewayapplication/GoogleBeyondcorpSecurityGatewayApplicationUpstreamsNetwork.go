package googlebeyondcorpsecuritygatewayapplication


type GoogleBeyondcorpSecurityGatewayApplicationUpstreamsNetwork struct {
	// Required. Network name is of the format: 'projects/{project}/global/networks/{network}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_beyondcorp_security_gateway_application#name GoogleBeyondcorpSecurityGatewayApplication#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

