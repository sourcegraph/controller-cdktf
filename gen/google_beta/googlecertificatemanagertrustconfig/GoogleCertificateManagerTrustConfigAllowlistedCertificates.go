package googlecertificatemanagertrustconfig


type GoogleCertificateManagerTrustConfigAllowlistedCertificates struct {
	// PEM certificate that is allowlisted.
	//
	// The certificate can be up to 5k bytes, and must be a parseable X.509 certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_certificate_manager_trust_config#pem_certificate GoogleCertificateManagerTrustConfig#pem_certificate}
	PemCertificate *string `field:"required" json:"pemCertificate" yaml:"pemCertificate"`
}

