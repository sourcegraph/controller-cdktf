package googleprivatecacertificatetemplate


type GooglePrivatecaCertificateTemplatePassthroughExtensions struct {
	// additional_extensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate_template#additional_extensions GooglePrivatecaCertificateTemplate#additional_extensions}
	AdditionalExtensions interface{} `field:"optional" json:"additionalExtensions" yaml:"additionalExtensions"`
	// Optional.
	//
	// A set of named X.509 extensions. Will be combined with additional_extensions to determine the full set of X.509 extensions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate_template#known_extensions GooglePrivatecaCertificateTemplate#known_extensions}
	KnownExtensions *[]*string `field:"optional" json:"knownExtensions" yaml:"knownExtensions"`
}

