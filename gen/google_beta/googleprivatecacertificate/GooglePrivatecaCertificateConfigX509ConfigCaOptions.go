package googleprivatecacertificate


type GooglePrivatecaCertificateConfigX509ConfigCaOptions struct {
	// When true, the "CA" in Basic Constraints extension will be set to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate#is_ca GooglePrivatecaCertificate#is_ca}
	IsCa interface{} `field:"optional" json:"isCa" yaml:"isCa"`
	// Refers to the "path length constraint" in Basic Constraints extension.
	//
	// For a CA certificate, this value describes the depth of
	// subordinate CA certificates that are allowed. If this value is less than 0, the request will fail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate#max_issuer_path_length GooglePrivatecaCertificate#max_issuer_path_length}
	MaxIssuerPathLength *float64 `field:"optional" json:"maxIssuerPathLength" yaml:"maxIssuerPathLength"`
	// When true, the "CA" in Basic Constraints extension will be set to false.
	//
	// If both 'is_ca' and 'non_ca' are unset, the extension will be omitted from the CA certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate#non_ca GooglePrivatecaCertificate#non_ca}
	NonCa interface{} `field:"optional" json:"nonCa" yaml:"nonCa"`
	// When true, the "path length constraint" in Basic Constraints extension will be set to 0.
	//
	// if both 'max_issuer_path_length' and 'zero_max_issuer_path_length' are unset,
	// the max path length will be omitted from the CA certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate#zero_max_issuer_path_length GooglePrivatecaCertificate#zero_max_issuer_path_length}
	ZeroMaxIssuerPathLength interface{} `field:"optional" json:"zeroMaxIssuerPathLength" yaml:"zeroMaxIssuerPathLength"`
}

