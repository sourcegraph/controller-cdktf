package googleapigeeapiproduct


type GoogleApigeeApiProductAttributes struct {
	// Key of the attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_api_product#name GoogleApigeeApiProduct#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Value of the attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_api_product#value GoogleApigeeApiProduct#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

