package googleprivatecacertificatetemplate


type GooglePrivatecaCertificateTemplatePredefinedValuesAdditionalExtensions struct {
	// object_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate_template#object_id GooglePrivatecaCertificateTemplate#object_id}
	ObjectId *GooglePrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId `field:"required" json:"objectId" yaml:"objectId"`
	// Required. The value of this X.509 extension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate_template#value GooglePrivatecaCertificateTemplate#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Optional.
	//
	// Indicates whether or not this extension is critical (i.e., if the client does not know how to handle this extension, the client should consider this to be an error).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate_template#critical GooglePrivatecaCertificateTemplate#critical}
	Critical interface{} `field:"optional" json:"critical" yaml:"critical"`
}

