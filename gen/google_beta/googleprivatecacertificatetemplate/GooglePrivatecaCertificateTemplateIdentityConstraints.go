package googleprivatecacertificatetemplate


type GooglePrivatecaCertificateTemplateIdentityConstraints struct {
	// Required.
	//
	// If this is true, the SubjectAltNames extension may be copied from a certificate request into the signed certificate. Otherwise, the requested SubjectAltNames will be discarded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate_template#allow_subject_alt_names_passthrough GooglePrivatecaCertificateTemplate#allow_subject_alt_names_passthrough}
	AllowSubjectAltNamesPassthrough interface{} `field:"required" json:"allowSubjectAltNamesPassthrough" yaml:"allowSubjectAltNamesPassthrough"`
	// Required.
	//
	// If this is true, the Subject field may be copied from a certificate request into the signed certificate. Otherwise, the requested Subject will be discarded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate_template#allow_subject_passthrough GooglePrivatecaCertificateTemplate#allow_subject_passthrough}
	AllowSubjectPassthrough interface{} `field:"required" json:"allowSubjectPassthrough" yaml:"allowSubjectPassthrough"`
	// cel_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate_template#cel_expression GooglePrivatecaCertificateTemplate#cel_expression}
	CelExpression *GooglePrivatecaCertificateTemplateIdentityConstraintsCelExpression `field:"optional" json:"celExpression" yaml:"celExpression"`
}

