package googlecertificatemanagertrustconfig


type GoogleCertificateManagerTrustConfigTrustStoresIntermediateCas struct {
	// PEM intermediate certificate used for building up paths for validation.
	//
	// Each certificate provided in PEM format may occupy up to 5kB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_certificate_manager_trust_config#pem_certificate GoogleCertificateManagerTrustConfig#pem_certificate}
	PemCertificate *string `field:"optional" json:"pemCertificate" yaml:"pemCertificate"`
}

