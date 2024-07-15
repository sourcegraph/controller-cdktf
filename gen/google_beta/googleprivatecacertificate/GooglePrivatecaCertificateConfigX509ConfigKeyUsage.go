package googleprivatecacertificate


type GooglePrivatecaCertificateConfigX509ConfigKeyUsage struct {
	// base_key_usage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate#base_key_usage GooglePrivatecaCertificate#base_key_usage}
	BaseKeyUsage *GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsage `field:"required" json:"baseKeyUsage" yaml:"baseKeyUsage"`
	// extended_key_usage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate#extended_key_usage GooglePrivatecaCertificate#extended_key_usage}
	ExtendedKeyUsage *GooglePrivatecaCertificateConfigX509ConfigKeyUsageExtendedKeyUsage `field:"required" json:"extendedKeyUsage" yaml:"extendedKeyUsage"`
	// unknown_extended_key_usages block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate#unknown_extended_key_usages GooglePrivatecaCertificate#unknown_extended_key_usages}
	UnknownExtendedKeyUsages interface{} `field:"optional" json:"unknownExtendedKeyUsages" yaml:"unknownExtendedKeyUsages"`
}

