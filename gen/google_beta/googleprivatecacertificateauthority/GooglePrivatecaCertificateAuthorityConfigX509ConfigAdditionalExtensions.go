package googleprivatecacertificateauthority


type GooglePrivatecaCertificateAuthorityConfigX509ConfigAdditionalExtensions struct {
	// Indicates whether or not this extension is critical (i.e., if the client does not know how to handle this extension, the client should consider this to be an error).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate_authority#critical GooglePrivatecaCertificateAuthority#critical}
	Critical interface{} `field:"required" json:"critical" yaml:"critical"`
	// object_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate_authority#object_id GooglePrivatecaCertificateAuthority#object_id}
	ObjectId *GooglePrivatecaCertificateAuthorityConfigX509ConfigAdditionalExtensionsObjectId `field:"required" json:"objectId" yaml:"objectId"`
	// The value of this X.509 extension. A base64-encoded string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate_authority#value GooglePrivatecaCertificateAuthority#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

