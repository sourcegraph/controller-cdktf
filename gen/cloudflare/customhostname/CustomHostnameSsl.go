package customhostname


type CustomHostnameSsl struct {
	// A ubiquitous bundle has the highest probability of being verified everywhere, even by clients using outdated or unusual trust stores.
	//
	// An optimal bundle uses the shortest chain and newest intermediates. And the force bundle verifies the chain, but does not otherwise modify it. Available values: `ubiquitous`, `optimal`, `force`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/custom_hostname#bundle_method CustomHostname#bundle_method}
	BundleMethod *string `field:"optional" json:"bundleMethod" yaml:"bundleMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/custom_hostname#certificate_authority CustomHostname#certificate_authority}.
	CertificateAuthority *string `field:"optional" json:"certificateAuthority" yaml:"certificateAuthority"`
	// If a custom uploaded certificate is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/custom_hostname#custom_certificate CustomHostname#custom_certificate}
	CustomCertificate *string `field:"optional" json:"customCertificate" yaml:"customCertificate"`
	// The key for a custom uploaded certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/custom_hostname#custom_key CustomHostname#custom_key}
	CustomKey *string `field:"optional" json:"customKey" yaml:"customKey"`
	// Domain control validation (DCV) method used for this hostname. Available values: `http`, `txt`, `email`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/custom_hostname#method CustomHostname#method}
	Method *string `field:"optional" json:"method" yaml:"method"`
	// settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/custom_hostname#settings CustomHostname#settings}
	Settings interface{} `field:"optional" json:"settings" yaml:"settings"`
	// Level of validation to be used for this hostname. Available values: `dv`. Defaults to `dv`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/custom_hostname#type CustomHostname#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Indicates whether the certificate covers a wildcard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/custom_hostname#wildcard CustomHostname#wildcard}
	Wildcard interface{} `field:"optional" json:"wildcard" yaml:"wildcard"`
}

