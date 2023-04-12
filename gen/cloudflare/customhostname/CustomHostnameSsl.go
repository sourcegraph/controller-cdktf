package customhostname


type CustomHostnameSsl struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/custom_hostname#certificate_authority CustomHostname#certificate_authority}.
	CertificateAuthority *string `field:"optional" json:"certificateAuthority" yaml:"certificateAuthority"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/custom_hostname#custom_certificate CustomHostname#custom_certificate}.
	CustomCertificate *string `field:"optional" json:"customCertificate" yaml:"customCertificate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/custom_hostname#custom_key CustomHostname#custom_key}.
	CustomKey *string `field:"optional" json:"customKey" yaml:"customKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/custom_hostname#method CustomHostname#method}.
	Method *string `field:"optional" json:"method" yaml:"method"`
	// settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/custom_hostname#settings CustomHostname#settings}
	Settings interface{} `field:"optional" json:"settings" yaml:"settings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/custom_hostname#type CustomHostname#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/custom_hostname#wildcard CustomHostname#wildcard}.
	Wildcard interface{} `field:"optional" json:"wildcard" yaml:"wildcard"`
}

