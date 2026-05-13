package googlebeyondcorpsecuritygateway


type GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeaders struct {
	// device_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_beyondcorp_security_gateway#device_info GoogleBeyondcorpSecurityGateway#device_info}
	DeviceInfo *GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersDeviceInfo `field:"optional" json:"deviceInfo" yaml:"deviceInfo"`
	// group_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_beyondcorp_security_gateway#group_info GoogleBeyondcorpSecurityGateway#group_info}
	GroupInfo *GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersGroupInfo `field:"optional" json:"groupInfo" yaml:"groupInfo"`
	// Default output type for all enabled headers. Possible values: ["PROTOBUF", "JSON", "NONE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_beyondcorp_security_gateway#output_type GoogleBeyondcorpSecurityGateway#output_type}
	OutputType *string `field:"optional" json:"outputType" yaml:"outputType"`
	// user_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_beyondcorp_security_gateway#user_info GoogleBeyondcorpSecurityGateway#user_info}
	UserInfo *GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersUserInfo `field:"optional" json:"userInfo" yaml:"userInfo"`
}

