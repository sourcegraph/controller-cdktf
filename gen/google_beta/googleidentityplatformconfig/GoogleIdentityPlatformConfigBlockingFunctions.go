package googleidentityplatformconfig


type GoogleIdentityPlatformConfigBlockingFunctions struct {
	// triggers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#triggers GoogleIdentityPlatformConfig#triggers}
	Triggers interface{} `field:"required" json:"triggers" yaml:"triggers"`
	// forward_inbound_credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#forward_inbound_credentials GoogleIdentityPlatformConfig#forward_inbound_credentials}
	ForwardInboundCredentials *GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentials `field:"optional" json:"forwardInboundCredentials" yaml:"forwardInboundCredentials"`
}

