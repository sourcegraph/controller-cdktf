package googleapigatewayapiconfig


type GoogleApiGatewayApiConfigGrpcServices struct {
	// file_descriptor_set block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_api_gateway_api_config#file_descriptor_set GoogleApiGatewayApiConfigA#file_descriptor_set}
	FileDescriptorSet *GoogleApiGatewayApiConfigGrpcServicesFileDescriptorSet `field:"required" json:"fileDescriptorSet" yaml:"fileDescriptorSet"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_api_gateway_api_config#source GoogleApiGatewayApiConfigA#source}
	Source interface{} `field:"optional" json:"source" yaml:"source"`
}

