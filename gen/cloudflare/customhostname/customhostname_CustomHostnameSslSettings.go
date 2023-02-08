package customhostname


type CustomHostnameSslSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/custom_hostname#ciphers CustomHostname#ciphers}.
	Ciphers *[]*string `field:"optional" json:"ciphers" yaml:"ciphers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/custom_hostname#early_hints CustomHostname#early_hints}.
	EarlyHints *string `field:"optional" json:"earlyHints" yaml:"earlyHints"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/custom_hostname#http2 CustomHostname#http2}.
	Http2 *string `field:"optional" json:"http2" yaml:"http2"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/custom_hostname#min_tls_version CustomHostname#min_tls_version}.
	MinTlsVersion *string `field:"optional" json:"minTlsVersion" yaml:"minTlsVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/custom_hostname#tls13 CustomHostname#tls13}.
	Tls13 *string `field:"optional" json:"tls13" yaml:"tls13"`
}

