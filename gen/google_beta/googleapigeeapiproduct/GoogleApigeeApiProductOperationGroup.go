package googleapigeeapiproduct


type GoogleApigeeApiProductOperationGroup struct {
	// operation_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_api_product#operation_configs GoogleApigeeApiProduct#operation_configs}
	OperationConfigs interface{} `field:"optional" json:"operationConfigs" yaml:"operationConfigs"`
	// Flag that specifes whether the configuration is for Apigee API proxy or a remote service.
	//
	// Valid values include proxy or remoteservice. Defaults to proxy. Set to proxy when Apigee API proxies are associated with the API product. Set to remoteservice when non-Apigee proxies like Istio-Envoy are associated with the API product. Possible values: ["proxy", "remoteservice"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_api_product#operation_config_type GoogleApigeeApiProduct#operation_config_type}
	OperationConfigType *string `field:"optional" json:"operationConfigType" yaml:"operationConfigType"`
}

