package googleprivatecacertificatetemplate


type GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsage struct {
	// base_key_usage block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_privateca_certificate_template#base_key_usage GooglePrivatecaCertificateTemplate#base_key_usage}
	BaseKeyUsage *GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage `field:"optional" json:"baseKeyUsage" yaml:"baseKeyUsage"`
	// extended_key_usage block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_privateca_certificate_template#extended_key_usage GooglePrivatecaCertificateTemplate#extended_key_usage}
	ExtendedKeyUsage *GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage `field:"optional" json:"extendedKeyUsage" yaml:"extendedKeyUsage"`
	// unknown_extended_key_usages block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_privateca_certificate_template#unknown_extended_key_usages GooglePrivatecaCertificateTemplate#unknown_extended_key_usages}
	UnknownExtendedKeyUsages interface{} `field:"optional" json:"unknownExtendedKeyUsages" yaml:"unknownExtendedKeyUsages"`
}

