package googlecomputemanagedsslcertificate


type GoogleComputeManagedSslCertificateManaged struct {
	// Domains for which a managed SSL certificate will be valid.
	//
	// Currently,
	// there can be up to 100 domains in this list.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_managed_ssl_certificate#domains GoogleComputeManagedSslCertificate#domains}
	Domains *[]*string `field:"required" json:"domains" yaml:"domains"`
}

