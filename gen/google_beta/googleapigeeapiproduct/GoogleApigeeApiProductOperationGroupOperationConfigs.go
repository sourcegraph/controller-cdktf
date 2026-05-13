package googleapigeeapiproduct


type GoogleApigeeApiProductOperationGroupOperationConfigs struct {
	// Required. Name of the API proxy or remote service with which the resources, methods, and quota are associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_api_product#api_source GoogleApigeeApiProduct#api_source}
	ApiSource *string `field:"optional" json:"apiSource" yaml:"apiSource"`
	// attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_api_product#attributes GoogleApigeeApiProduct#attributes}
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// operations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_api_product#operations GoogleApigeeApiProduct#operations}
	Operations interface{} `field:"optional" json:"operations" yaml:"operations"`
	// quota block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_api_product#quota GoogleApigeeApiProduct#quota}
	Quota *GoogleApigeeApiProductOperationGroupOperationConfigsQuota `field:"optional" json:"quota" yaml:"quota"`
}

