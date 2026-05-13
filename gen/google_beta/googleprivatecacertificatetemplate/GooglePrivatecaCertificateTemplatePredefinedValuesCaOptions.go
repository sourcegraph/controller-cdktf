package googleprivatecacertificatetemplate


type GooglePrivatecaCertificateTemplatePredefinedValuesCaOptions struct {
	// Optional.
	//
	// Refers to the "CA" X.509 extension, which is a boolean value. When this value is true, the "CA" in Basic Constraints extension will be set to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_privateca_certificate_template#is_ca GooglePrivatecaCertificateTemplate#is_ca}
	IsCa interface{} `field:"optional" json:"isCa" yaml:"isCa"`
	// Optional.
	//
	// Refers to the "path length constraint" in Basic Constraints extension. For a CA certificate, this value describes the depth of
	// subordinate CA certificates that are allowed. If this value is less than 0, the request will fail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_privateca_certificate_template#max_issuer_path_length GooglePrivatecaCertificateTemplate#max_issuer_path_length}
	MaxIssuerPathLength *float64 `field:"optional" json:"maxIssuerPathLength" yaml:"maxIssuerPathLength"`
	// Optional.
	//
	// When true, the "CA" in Basic Constraints extension will be set to null and omitted from the CA certificate.
	// If both 'is_ca' and 'null_ca' are unset, the "CA" in Basic Constraints extension will be set to false.
	// Note that the behavior when 'is_ca = false' for this resource is different from the behavior in the Certificate Authority, Certificate and CaPool resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_privateca_certificate_template#null_ca GooglePrivatecaCertificateTemplate#null_ca}
	NullCa interface{} `field:"optional" json:"nullCa" yaml:"nullCa"`
	// Optional.
	//
	// When true, the "path length constraint" in Basic Constraints extension will be set to 0.
	// if both 'max_issuer_path_length' and 'zero_max_issuer_path_length' are unset,
	// the max path length will be omitted from the CA certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_privateca_certificate_template#zero_max_issuer_path_length GooglePrivatecaCertificateTemplate#zero_max_issuer_path_length}
	ZeroMaxIssuerPathLength interface{} `field:"optional" json:"zeroMaxIssuerPathLength" yaml:"zeroMaxIssuerPathLength"`
}

