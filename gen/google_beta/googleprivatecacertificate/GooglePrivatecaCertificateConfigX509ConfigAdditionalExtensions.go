package googleprivatecacertificate


type GooglePrivatecaCertificateConfigX509ConfigAdditionalExtensions struct {
	// Indicates whether or not this extension is critical (i.e., if the client does not know how to handle this extension, the client should consider this to be an error).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate#critical GooglePrivatecaCertificate#critical}
	Critical interface{} `field:"required" json:"critical" yaml:"critical"`
	// object_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate#object_id GooglePrivatecaCertificate#object_id}
	ObjectId *GooglePrivatecaCertificateConfigX509ConfigAdditionalExtensionsObjectId `field:"required" json:"objectId" yaml:"objectId"`
	// The value of this X.509 extension. A base64-encoded string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate#value GooglePrivatecaCertificate#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

