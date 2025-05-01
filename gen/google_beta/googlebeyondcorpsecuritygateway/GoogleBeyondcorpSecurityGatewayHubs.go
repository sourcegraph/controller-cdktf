package googlebeyondcorpsecuritygateway


type GoogleBeyondcorpSecurityGatewayHubs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_beyondcorp_security_gateway#region GoogleBeyondcorpSecurityGateway#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
	// internet_gateway block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_beyondcorp_security_gateway#internet_gateway GoogleBeyondcorpSecurityGateway#internet_gateway}
	InternetGateway *GoogleBeyondcorpSecurityGatewayHubsInternetGateway `field:"optional" json:"internetGateway" yaml:"internetGateway"`
}

