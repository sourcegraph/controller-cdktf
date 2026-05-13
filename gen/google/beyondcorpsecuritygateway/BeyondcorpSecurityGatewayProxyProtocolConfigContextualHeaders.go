package beyondcorpsecuritygateway


type BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeaders struct {
	// device_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/beyondcorp_security_gateway#device_info BeyondcorpSecurityGateway#device_info}
	DeviceInfo *BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersDeviceInfo `field:"optional" json:"deviceInfo" yaml:"deviceInfo"`
	// group_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/beyondcorp_security_gateway#group_info BeyondcorpSecurityGateway#group_info}
	GroupInfo *BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersGroupInfo `field:"optional" json:"groupInfo" yaml:"groupInfo"`
	// Default output type for all enabled headers. Possible values: ["PROTOBUF", "JSON", "NONE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/beyondcorp_security_gateway#output_type BeyondcorpSecurityGateway#output_type}
	OutputType *string `field:"optional" json:"outputType" yaml:"outputType"`
	// user_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/beyondcorp_security_gateway#user_info BeyondcorpSecurityGateway#user_info}
	UserInfo *BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersUserInfo `field:"optional" json:"userInfo" yaml:"userInfo"`
}

