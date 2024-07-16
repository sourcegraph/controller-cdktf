package googlecertificatemanagercertificateissuanceconfig


type GoogleCertificateManagerCertificateIssuanceConfigCertificateAuthorityConfigCertificateAuthorityServiceConfig struct {
	// A CA pool resource used to issue a certificate.
	//
	// The CA pool string has a relative resource path following the form
	// "projects/{project}/locations/{location}/caPools/{caPool}".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_certificate_manager_certificate_issuance_config#ca_pool GoogleCertificateManagerCertificateIssuanceConfig#ca_pool}
	CaPool *string `field:"required" json:"caPool" yaml:"caPool"`
}

