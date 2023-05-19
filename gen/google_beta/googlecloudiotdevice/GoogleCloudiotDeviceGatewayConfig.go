package googlecloudiotdevice


type GoogleCloudiotDeviceGatewayConfig struct {
	// Indicates whether the device is a gateway. Possible values: ["ASSOCIATION_ONLY", "DEVICE_AUTH_TOKEN_ONLY", "ASSOCIATION_AND_DEVICE_AUTH_TOKEN"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_cloudiot_device#gateway_auth_method GoogleCloudiotDevice#gateway_auth_method}
	GatewayAuthMethod *string `field:"optional" json:"gatewayAuthMethod" yaml:"gatewayAuthMethod"`
	// Indicates whether the device is a gateway. Default value: "NON_GATEWAY" Possible values: ["GATEWAY", "NON_GATEWAY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_cloudiot_device#gateway_type GoogleCloudiotDevice#gateway_type}
	GatewayType *string `field:"optional" json:"gatewayType" yaml:"gatewayType"`
}

