package beyondcorpsecuritygatewayapplication


type BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocol struct {
	// The configuration for the proxy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/beyondcorp_security_gateway_application#allowed_client_headers BeyondcorpSecurityGatewayApplication#allowed_client_headers}
	AllowedClientHeaders *[]*string `field:"optional" json:"allowedClientHeaders" yaml:"allowedClientHeaders"`
	// Client IP configuration. The client IP address is included if true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/beyondcorp_security_gateway_application#client_ip BeyondcorpSecurityGatewayApplication#client_ip}
	ClientIp interface{} `field:"optional" json:"clientIp" yaml:"clientIp"`
	// contextual_headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/beyondcorp_security_gateway_application#contextual_headers BeyondcorpSecurityGatewayApplication#contextual_headers}
	ContextualHeaders *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeaders `field:"optional" json:"contextualHeaders" yaml:"contextualHeaders"`
	// Gateway identity configuration. Possible values: ["RESOURCE_NAME"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/beyondcorp_security_gateway_application#gateway_identity BeyondcorpSecurityGatewayApplication#gateway_identity}
	GatewayIdentity *string `field:"optional" json:"gatewayIdentity" yaml:"gatewayIdentity"`
	// Custom resource specific headers along with the values.
	//
	// The names should conform to RFC 9110:
	// > Field names SHOULD constrain themselves to alphanumeric characters, "-",
	//   and ".", and SHOULD begin with a letter.
	// > Field values SHOULD contain only ASCII printable characters and tab.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/beyondcorp_security_gateway_application#metadata_headers BeyondcorpSecurityGatewayApplication#metadata_headers}
	MetadataHeaders *map[string]*string `field:"optional" json:"metadataHeaders" yaml:"metadataHeaders"`
}

