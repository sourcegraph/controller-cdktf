package googlebeyondcorpsecuritygateway


type GoogleBeyondcorpSecurityGatewayProxyProtocolConfig struct {
	// The configuration for the proxy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_beyondcorp_security_gateway#allowed_client_headers GoogleBeyondcorpSecurityGateway#allowed_client_headers}
	AllowedClientHeaders *[]*string `field:"optional" json:"allowedClientHeaders" yaml:"allowedClientHeaders"`
	// Client IP configuration. The client IP address is included if true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_beyondcorp_security_gateway#client_ip GoogleBeyondcorpSecurityGateway#client_ip}
	ClientIp interface{} `field:"optional" json:"clientIp" yaml:"clientIp"`
	// contextual_headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_beyondcorp_security_gateway#contextual_headers GoogleBeyondcorpSecurityGateway#contextual_headers}
	ContextualHeaders *GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeaders `field:"optional" json:"contextualHeaders" yaml:"contextualHeaders"`
	// Gateway identity configuration. Possible values: ["RESOURCE_NAME"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_beyondcorp_security_gateway#gateway_identity GoogleBeyondcorpSecurityGateway#gateway_identity}
	GatewayIdentity *string `field:"optional" json:"gatewayIdentity" yaml:"gatewayIdentity"`
	// Custom resource specific headers along with the values.
	//
	// The names should conform to RFC 9110:
	// > Field names SHOULD constrain themselves to alphanumeric characters, "-",
	//   and ".", and SHOULD begin with a letter.
	// > Field values SHOULD contain only ASCII printable characters and tab.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_beyondcorp_security_gateway#metadata_headers GoogleBeyondcorpSecurityGateway#metadata_headers}
	MetadataHeaders *map[string]*string `field:"optional" json:"metadataHeaders" yaml:"metadataHeaders"`
}

