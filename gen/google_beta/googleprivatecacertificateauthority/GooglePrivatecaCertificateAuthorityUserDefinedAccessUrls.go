package googleprivatecacertificateauthority


type GooglePrivatecaCertificateAuthorityUserDefinedAccessUrls struct {
	// A list of URLs where this CertificateAuthority's CA certificate is published that is specified by users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_privateca_certificate_authority#aia_issuing_certificate_urls GooglePrivatecaCertificateAuthority#aia_issuing_certificate_urls}
	AiaIssuingCertificateUrls *[]*string `field:"optional" json:"aiaIssuingCertificateUrls" yaml:"aiaIssuingCertificateUrls"`
	// A list of URLs where this CertificateAuthority's CRLs are published that is specified by users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_privateca_certificate_authority#crl_access_urls GooglePrivatecaCertificateAuthority#crl_access_urls}
	CrlAccessUrls *[]*string `field:"optional" json:"crlAccessUrls" yaml:"crlAccessUrls"`
}

