package customhostname


type CustomHostnameSslSettings struct {
	// List of SSL/TLS ciphers to associate with this certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/custom_hostname#ciphers CustomHostname#ciphers}
	Ciphers *[]*string `field:"optional" json:"ciphers" yaml:"ciphers"`
	// Whether early hints should be supported. Available values: `on`, `off`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/custom_hostname#early_hints CustomHostname#early_hints}
	EarlyHints *string `field:"optional" json:"earlyHints" yaml:"earlyHints"`
	// Whether HTTP2 should be supported. Available values: `on`, `off`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/custom_hostname#http2 CustomHostname#http2}
	Http2 *string `field:"optional" json:"http2" yaml:"http2"`
	// Lowest version of TLS this certificate should support. Available values: `1.0`, `1.1`, `1.2`, `1.3`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/custom_hostname#min_tls_version CustomHostname#min_tls_version}
	MinTlsVersion *string `field:"optional" json:"minTlsVersion" yaml:"minTlsVersion"`
	// Whether TLSv1.3 should be supported. Available values: `on`, `off`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/custom_hostname#tls13 CustomHostname#tls13}
	Tls13 *string `field:"optional" json:"tls13" yaml:"tls13"`
}

