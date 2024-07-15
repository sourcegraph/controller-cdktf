package googleprivatecacertificateauthority


type GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsage struct {
	// base_key_usage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate_authority#base_key_usage GooglePrivatecaCertificateAuthority#base_key_usage}
	BaseKeyUsage *GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage `field:"required" json:"baseKeyUsage" yaml:"baseKeyUsage"`
	// extended_key_usage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate_authority#extended_key_usage GooglePrivatecaCertificateAuthority#extended_key_usage}
	ExtendedKeyUsage *GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage `field:"required" json:"extendedKeyUsage" yaml:"extendedKeyUsage"`
	// unknown_extended_key_usages block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate_authority#unknown_extended_key_usages GooglePrivatecaCertificateAuthority#unknown_extended_key_usages}
	UnknownExtendedKeyUsages interface{} `field:"optional" json:"unknownExtendedKeyUsages" yaml:"unknownExtendedKeyUsages"`
}

