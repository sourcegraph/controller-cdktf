package googleprivatecacertificateauthority


type GooglePrivatecaCertificateAuthorityConfigX509Config struct {
	// ca_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate_authority#ca_options GooglePrivatecaCertificateAuthority#ca_options}
	CaOptions *GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptions `field:"required" json:"caOptions" yaml:"caOptions"`
	// key_usage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate_authority#key_usage GooglePrivatecaCertificateAuthority#key_usage}
	KeyUsage *GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsage `field:"required" json:"keyUsage" yaml:"keyUsage"`
	// additional_extensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate_authority#additional_extensions GooglePrivatecaCertificateAuthority#additional_extensions}
	AdditionalExtensions interface{} `field:"optional" json:"additionalExtensions" yaml:"additionalExtensions"`
	// Describes Online Certificate Status Protocol (OCSP) endpoint addresses that appear in the "Authority Information Access" extension in the certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate_authority#aia_ocsp_servers GooglePrivatecaCertificateAuthority#aia_ocsp_servers}
	AiaOcspServers *[]*string `field:"optional" json:"aiaOcspServers" yaml:"aiaOcspServers"`
	// name_constraints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate_authority#name_constraints GooglePrivatecaCertificateAuthority#name_constraints}
	NameConstraints *GooglePrivatecaCertificateAuthorityConfigX509ConfigNameConstraints `field:"optional" json:"nameConstraints" yaml:"nameConstraints"`
	// policy_ids block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate_authority#policy_ids GooglePrivatecaCertificateAuthority#policy_ids}
	PolicyIds interface{} `field:"optional" json:"policyIds" yaml:"policyIds"`
}

