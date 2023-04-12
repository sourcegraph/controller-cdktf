package googlecertificatemanagercertificate


type GoogleCertificateManagerCertificateSelfManaged struct {
	// The certificate chain in PEM-encoded form.
	//
	// Leaf certificate comes first, followed by intermediate ones if any.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_certificate_manager_certificate#certificate_pem GoogleCertificateManagerCertificate#certificate_pem}
	CertificatePem *string `field:"required" json:"certificatePem" yaml:"certificatePem"`
	// The private key of the leaf certificate in PEM-encoded form.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_certificate_manager_certificate#private_key_pem GoogleCertificateManagerCertificate#private_key_pem}
	PrivateKeyPem *string `field:"required" json:"privateKeyPem" yaml:"privateKeyPem"`
}

