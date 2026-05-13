package googlebeyondcorpsecuritygatewayapplication


type GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocol struct {
	// The configuration for the proxy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_beyondcorp_security_gateway_application#allowed_client_headers GoogleBeyondcorpSecurityGatewayApplication#allowed_client_headers}
	AllowedClientHeaders *[]*string `field:"optional" json:"allowedClientHeaders" yaml:"allowedClientHeaders"`
	// Client IP configuration. The client IP address is included if true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_beyondcorp_security_gateway_application#client_ip GoogleBeyondcorpSecurityGatewayApplication#client_ip}
	ClientIp interface{} `field:"optional" json:"clientIp" yaml:"clientIp"`
	// contextual_headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_beyondcorp_security_gateway_application#contextual_headers GoogleBeyondcorpSecurityGatewayApplication#contextual_headers}
	ContextualHeaders *GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeaders `field:"optional" json:"contextualHeaders" yaml:"contextualHeaders"`
	// Gateway identity configuration. Possible values: ["RESOURCE_NAME"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_beyondcorp_security_gateway_application#gateway_identity GoogleBeyondcorpSecurityGatewayApplication#gateway_identity}
	GatewayIdentity *string `field:"optional" json:"gatewayIdentity" yaml:"gatewayIdentity"`
	// Custom resource specific headers along with the values.
	//
	// The names should conform to RFC 9110:
	// > Field names SHOULD constrain themselves to alphanumeric characters, "-",
	//   and ".", and SHOULD begin with a letter.
	// > Field values SHOULD contain only ASCII printable characters and tab.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_beyondcorp_security_gateway_application#metadata_headers GoogleBeyondcorpSecurityGatewayApplication#metadata_headers}
	MetadataHeaders *map[string]*string `field:"optional" json:"metadataHeaders" yaml:"metadataHeaders"`
}

