package googlebeyondcorpsecuritygatewayapplication


type GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeaders struct {
	// device_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_beyondcorp_security_gateway_application#device_info GoogleBeyondcorpSecurityGatewayApplication#device_info}
	DeviceInfo *GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersDeviceInfo `field:"optional" json:"deviceInfo" yaml:"deviceInfo"`
	// group_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_beyondcorp_security_gateway_application#group_info GoogleBeyondcorpSecurityGatewayApplication#group_info}
	GroupInfo *GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersGroupInfo `field:"optional" json:"groupInfo" yaml:"groupInfo"`
	// Default output type for all enabled headers. Possible values: ["PROTOBUF", "JSON", "NONE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_beyondcorp_security_gateway_application#output_type GoogleBeyondcorpSecurityGatewayApplication#output_type}
	OutputType *string `field:"optional" json:"outputType" yaml:"outputType"`
	// user_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_beyondcorp_security_gateway_application#user_info GoogleBeyondcorpSecurityGatewayApplication#user_info}
	UserInfo *GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfo `field:"optional" json:"userInfo" yaml:"userInfo"`
}

