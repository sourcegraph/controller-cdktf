package googleprivatecacertificatetemplate


type GooglePrivatecaCertificateTemplatePredefinedValues struct {
	// additional_extensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate_template#additional_extensions GooglePrivatecaCertificateTemplate#additional_extensions}
	AdditionalExtensions interface{} `field:"optional" json:"additionalExtensions" yaml:"additionalExtensions"`
	// Optional.
	//
	// Describes Online Certificate Status Protocol (OCSP) endpoint addresses that appear in the "Authority Information Access" extension in the certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate_template#aia_ocsp_servers GooglePrivatecaCertificateTemplate#aia_ocsp_servers}
	AiaOcspServers *[]*string `field:"optional" json:"aiaOcspServers" yaml:"aiaOcspServers"`
	// ca_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate_template#ca_options GooglePrivatecaCertificateTemplate#ca_options}
	CaOptions *GooglePrivatecaCertificateTemplatePredefinedValuesCaOptions `field:"optional" json:"caOptions" yaml:"caOptions"`
	// key_usage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate_template#key_usage GooglePrivatecaCertificateTemplate#key_usage}
	KeyUsage *GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsage `field:"optional" json:"keyUsage" yaml:"keyUsage"`
	// policy_ids block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate_template#policy_ids GooglePrivatecaCertificateTemplate#policy_ids}
	PolicyIds interface{} `field:"optional" json:"policyIds" yaml:"policyIds"`
}

