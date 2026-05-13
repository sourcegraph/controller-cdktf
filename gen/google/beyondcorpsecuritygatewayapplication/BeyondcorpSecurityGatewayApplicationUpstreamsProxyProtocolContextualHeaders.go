package beyondcorpsecuritygatewayapplication


type BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeaders struct {
	// device_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/beyondcorp_security_gateway_application#device_info BeyondcorpSecurityGatewayApplication#device_info}
	DeviceInfo *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersDeviceInfo `field:"optional" json:"deviceInfo" yaml:"deviceInfo"`
	// group_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/beyondcorp_security_gateway_application#group_info BeyondcorpSecurityGatewayApplication#group_info}
	GroupInfo *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersGroupInfo `field:"optional" json:"groupInfo" yaml:"groupInfo"`
	// Default output type for all enabled headers. Possible values: ["PROTOBUF", "JSON", "NONE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/beyondcorp_security_gateway_application#output_type BeyondcorpSecurityGatewayApplication#output_type}
	OutputType *string `field:"optional" json:"outputType" yaml:"outputType"`
	// user_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/beyondcorp_security_gateway_application#user_info BeyondcorpSecurityGatewayApplication#user_info}
	UserInfo *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfo `field:"optional" json:"userInfo" yaml:"userInfo"`
}

