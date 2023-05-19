package customssl


type CustomSslCustomSslOptions struct {
	// Method of building intermediate certificate chain.
	//
	// A ubiquitous bundle has the highest probability of being verified everywhere, even by clients using outdated or unusual trust stores. An optimal bundle uses the shortest chain and newest intermediates. And the force bundle verifies the chain, but does not otherwise modify it. Available values: `ubiquitous`, `optimal`, `force`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/custom_ssl#bundle_method CustomSsl#bundle_method}
	BundleMethod *string `field:"optional" json:"bundleMethod" yaml:"bundleMethod"`
	// Certificate certificate and the intermediate(s).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/custom_ssl#certificate CustomSsl#certificate}
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// Specifies the region where your private key can be held locally. Available values: `us`, `eu`, `highest_security`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/custom_ssl#geo_restrictions CustomSsl#geo_restrictions}
	GeoRestrictions *string `field:"optional" json:"geoRestrictions" yaml:"geoRestrictions"`
	// Certificate's private key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/custom_ssl#private_key CustomSsl#private_key}
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// Whether to enable support for legacy clients which do not include SNI in the TLS handshake.
	//
	// Available values: `legacy_custom`, `sni_custom`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/custom_ssl#type CustomSsl#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

