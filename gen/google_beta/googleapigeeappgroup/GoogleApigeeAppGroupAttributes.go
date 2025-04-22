package googleapigeeappgroup


type GoogleApigeeAppGroupAttributes struct {
	// Key of the attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_apigee_app_group#name GoogleApigeeAppGroup#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Value of the attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_apigee_app_group#value GoogleApigeeAppGroup#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

