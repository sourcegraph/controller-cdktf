package appmeshvirtualgateway


type AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNames struct {
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/appmesh_virtual_gateway#match AppmeshVirtualGateway#match}
	Match *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatch `field:"required" json:"match" yaml:"match"`
}

