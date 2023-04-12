package googleprivatecacertificateauthority


type GooglePrivatecaCertificateAuthoritySubordinateConfigPemIssuerChain struct {
	// Expected to be in leaf-to-root order according to RFC 5246.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_privateca_certificate_authority#pem_certificates GooglePrivatecaCertificateAuthority#pem_certificates}
	PemCertificates *[]*string `field:"optional" json:"pemCertificates" yaml:"pemCertificates"`
}

