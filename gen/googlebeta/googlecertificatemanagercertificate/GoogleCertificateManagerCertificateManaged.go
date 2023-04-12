package googlecertificatemanagercertificate


type GoogleCertificateManagerCertificateManaged struct {
	// Authorizations that will be used for performing domain authorization.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_certificate_manager_certificate#dns_authorizations GoogleCertificateManagerCertificate#dns_authorizations}
	DnsAuthorizations *[]*string `field:"optional" json:"dnsAuthorizations" yaml:"dnsAuthorizations"`
	// The domains for which a managed SSL certificate will be generated. Wildcard domains are only supported with DNS challenge resolution.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_certificate_manager_certificate#domains GoogleCertificateManagerCertificate#domains}
	Domains *[]*string `field:"optional" json:"domains" yaml:"domains"`
}

